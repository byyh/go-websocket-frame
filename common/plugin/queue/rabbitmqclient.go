package queue

import (
	"errors"
	"sync"
	"time"

	"go-websocket-frame/common/global/plugin/log"

	"github.com/streadway/amqp"
)

type InitParam struct {
	AmqpUrl         string `json:"amqp_url"`
	QueueName       string `json:"queue_name"`
	ExchangeName    string `json:"exchange_name"`
	ExchangeType    string `json:"exchange_type"`
	ExchangeDurable bool   `json:"exchange_durable"`
	RoutingKey      string `json:"routing_key"`
	Durable         bool   `json:"durable"`
}

type RabbitMqClient struct {
	conn    *amqp.Connection
	Channel *amqp.Channel

	QueueName       string
	ExchangeName    string
	ExchangeType    string
	ExchangeDurable bool
	RoutingKey      string
	Durable         bool
	MqUrl           string

	notifyCloseConn    chan *amqp.Error
	notifyCloseChannel chan *amqp.Error
	notifyConfirm      chan amqp.Confirmation

	mustExitEvent bool
	forwait       chan bool

	count            int
	consumeFailedMap sync.Map // 消费失败记录

}

const (
	// 单个消费者最多处理多少个消息后退出
	MaxConsumerCount = 1000000

	// 消费处理失败暂停的毫秒数量
	MaxConsumerFailedWaitTimes = 100

	// 单条记录消费失败重试次数
	MaxConsumerFailedRepeatCount = 3
)

var ()

func NewRabbitMqClient(param *InitParam) (*RabbitMqClient, error) {
	var err error

	rmq := &RabbitMqClient{
		MqUrl:           param.AmqpUrl,
		QueueName:       param.QueueName,
		ExchangeName:    param.ExchangeName,
		ExchangeType:    param.ExchangeType,
		ExchangeDurable: param.ExchangeDurable,
		RoutingKey:      param.RoutingKey,
		Durable:         param.Durable,
	}

	rmq.forwait = make(chan bool)
	rmq.notifyCloseConn = make(chan *amqp.Error)
	rmq.notifyCloseChannel = make(chan *amqp.Error)

	rmq.CheckParameter()

	rmq.conn, err = amqp.Dial(rmq.MqUrl)
	if nil != err {
		log.Error("RabbitMqClient-err:failed to connect tp rabbitmq,", err)
		return nil, err
	}

	return rmq, nil
}

func (this *RabbitMqClient) CheckErr(err error, msg string) {
	if nil == err {
		return
	}

	log.Error("RabbitMqClient-err,", msg, err)
	panic(err)
}

// 初始化渠道并绑定
func (this *RabbitMqClient) NewChannelAndBind() error {
	if this.Channel == nil {
		this.NewChannel()
	}

	return this.Bind()
}

// 一个连接只有一个通道，避免tcp堵塞
func (this *RabbitMqClient) NewChannel() error {
	var err error

	this.Channel, err = this.conn.Channel()
	if nil != err {
		log.Error("RabbitMq channel err,", err)
		return err
	}

	return nil
}

func (this *RabbitMqClient) Close() {
	this.Channel.Close()
	this.conn.Close()

	this.Channel = nil
}

func (this *RabbitMqClient) ConsumeNoAck(handler ConsumerHandler) error {
	if this.Channel == nil {
		if err := this.NewChannelAndBind(); nil != err {
			return err
		}
	}

	this.ConnectCheckTimer()

	msgs, err := this.Channel.Consume(this.QueueName, "", true /*true*/, false, false, false, nil)
	if nil != err {
		log.Error("Rabbitmq create Channel.consume error ", err)
		return err
	}

	go func(msgs <-chan amqp.Delivery) {
		tms := time.Tick(time.Second)

		for {
			select {
			case d := <-msgs:
				this.handleMsg(d, true, handler)

				if this.mustExitEvent {
					log.Info("recv mustExitEvent")
					this.forwait <- true
					return
				}

				this.count++
				if MaxConsumerCount < this.count {
					log.Info("handle msg count gt ", MaxConsumerCount, ",ready quit")
					this.forwait <- true
					return
				}
			case <-tms:
				if this.mustExitEvent {
					log.Info("recv mustExitEvent-time")
					return
				}
			}
		}
	}(msgs)

	log.Info(" [*] Waiting for messages... ")

	<-this.forwait

	return nil
}

// 需要应答的消费
func (this *RabbitMqClient) ConsumeMustAck(handler ConsumerHandler) error {
	if this.Channel == nil {
		if err := this.NewChannelAndBind(); nil != err {
			return err
		}
	}

	this.ConnectCheckTimer()

	err := this.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if nil != err {
		log.Error("Rabbitmq channel.Qos error ", err)
		return err
	}

	msgs, err := this.Channel.Consume(this.QueueName, "", false /*autoAck*/, false, false, false, nil)
	if nil != err {
		log.Error("Rabbitmq create Channel.consume error ", err)
		return err
	}

	go func(msgs <-chan amqp.Delivery) {
		tms := time.Tick(time.Second)
		for {
			select {
			case d := <-msgs:
				this.handleMsg(d, false, handler)

				if this.mustExitEvent {
					log.Info("recv mustExitEvent")
					this.forwait <- true
					return
				}

				this.count++
				if MaxConsumerCount < this.count {
					log.Info("handle msg count gt ", MaxConsumerCount, ",ready quit")
					this.forwait <- true
					return
				}
			case <-tms:
				if this.mustExitEvent {
					log.Info("recv mustExitEvent-time")
					return
				}
			}
		}
	}(msgs)

	log.Info(" [*] Waiting for messages...")

	<-this.forwait

	return nil
}

// 自定义消费参数的消费
func (this *RabbitMqClient) ConsumeSelfDefine(queueName string, autoAck bool, handler ConsumerHandler) error {
	if this.Channel == nil {
		if err := this.NewChannelAndBind(); nil != err {
			return err
		}
	}

	this.ConnectCheckTimer()

	err := this.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if nil != err {
		log.Error("Rabbitmq channel.Qos error ", err)
		return err
	}

	msgs, err := this.Channel.Consume(queueName, queueName, autoAck, false, false, false, nil)
	if nil != err {
		log.Error("Rabbitmq create Channel.consume error ", err)
		return err
	}

	go func(msgs <-chan amqp.Delivery) {
		tms := time.Tick(time.Second)
		for {
			select {
			case d := <-msgs:
				this.handleMsg(d, autoAck, handler)

				if this.mustExitEvent {
					log.Info("recv mustExitEvent")
					this.forwait <- true
					return
				}

				this.count++
				if MaxConsumerCount < this.count {
					log.Info("handle msg count gt ", MaxConsumerCount, ",ready quit")
					this.forwait <- true
					return
				}
			case <-tms:
				if this.mustExitEvent {
					log.Info("recv mustExitEvent-time")
					return
				}
			}
		}
	}(msgs)

	log.Info(" [*] Waiting for messages...")

	<-this.forwait
	log.Info("Waiting for messages forwait ok")

	return nil
}

func (this *RabbitMqClient) handleMsg(d amqp.Delivery, autoAck bool, handler ConsumerHandler) {
	log.Info(",no: ", this.count, "接受到消息: ", d, ", ", string(d.Body), ", ", this.QueueName, ", ", this.ExchangeName)
	if 0 == len(d.Body) {
		log.Info("空消息")
		return
	}
	err := handler(d.Body)
	if nil != err {
		log.Error("rabbitmq.consume handle error,", d.DeliveryTag, ", 消费内容=", d, ",", err)
		if !autoAck {
			if v, ok := this.consumeFailedMap.LoadOrStore(d.DeliveryTag, 1); ok {
				if v.(uint64) > MaxConsumerFailedRepeatCount {
					// 不再消费该消息
					log.Error("rabbitmq.consume repeat failed count max:", string(d.Body))
					this.consumeFailedMap.Delete(d.DeliveryTag)

					this.Channel.Ack(d.DeliveryTag, false) // 应答

					return
				}

				this.consumeFailedMap.Store(d.DeliveryTag, v.(int)+1)
			}

			this.Channel.Nack(d.DeliveryTag, false, true)

			time.Sleep(time.Duration(MaxConsumerFailedWaitTimes) * time.Millisecond)

			return
		}
	}

	//time.Sleep(time.Duration(500) * time.Millisecond)
	if !autoAck {
		this.Channel.Ack(d.DeliveryTag, false)
	}
	//log.Info( ",ack , wait 3 s, ", d.DeliveryTag)

	//time.Sleep(time.Duration(500) * time.Millisecond)

}

func (this *RabbitMqClient) ConnectCheckTimer() {
	c := this.conn.NotifyClose(this.notifyCloseConn)

	go func() {
		select {
		case e, ok := <-c:
			var msg string
			if ok {
				msg = e.Reason
			}
			log.Info("connect is close by server", msg)

			this.mustExitEvent = true
		}
	}()

	cl := this.Channel.NotifyClose(this.notifyCloseChannel)
	go func() {
		log.Info("wait..")
		select {
		case e, ok := <-cl:
			var msg string
			if ok {
				msg = e.Reason
			}
			log.Info("channel is close by server", msg)

			this.mustExitEvent = true
		}
	}()
}

func (this *RabbitMqClient) Publish(exchange, routeKey string, data []byte) error {
	if nil == this.Channel {
		if err := this.NewChannel(); nil != err {
			return err
		}
	}

	err := this.Channel.Publish(
		exchange, // exchange
		routeKey, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         data,
			DeliveryMode: amqp.Transient, // 1=non-persistent/Transient, 2=Persistent
			Priority:     0})             // 0-9

	return err
}

func (this *RabbitMqClient) CheckParameter() {
	if "" == this.MqUrl {
		panic(errors.New("MqUrl is not allow empty"))
	}
}

func (this *RabbitMqClient) Bind() error {
	if err := this.Channel.ExchangeDeclare(
		this.ExchangeName,    // name of the exchange
		this.ExchangeType,    // type
		this.ExchangeDurable, // durable持久化
		false,                // delete when complete
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	); nil != err {
		log.Error("Rabbitmq Channel.ExchangeDeclare error ", err)
		return err
	}

	queue, err := this.Channel.QueueDeclare(
		this.QueueName, // name of the queue
		this.Durable,   // durable 持久化
		false,          // delete when unused
		false,          // exclusive
		false,          // noWait
		nil,            // arguments
	)
	if nil != err {
		log.Error("Rabbitmq Channel.QueueDeclare error ", err)
		return err
	}

	if err := this.Channel.QueueBind(
		queue.Name,        // namethis.QueueName of the queue
		this.RoutingKey,   // bindingKey
		this.ExchangeName, // sourceExchange
		false,             // noWait
		nil,               // arguments
	); nil != err {
		log.Error("Rabbitmq Channel.QueueBind error ", err)
		return err
	}

	return nil
}

func (this *RabbitMqClient) SecurityExit() {
	this.mustExitEvent = true
	log.Info("exit-channel-----", this.QueueName, this.ExchangeName)
	//this.Channel.Cancel(this.QueueName, false)
	this.Channel.Close()
}
