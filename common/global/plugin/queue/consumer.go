package queue

import (
	"errors"
	"go-websocket-frame/common/global/plugin/config"
	"go-websocket-frame/common/plugin/queue"
	"go-websocket-frame/common/utils"
	"sync"
	"sync/atomic"
	"time"

	"go-websocket-frame/common/global/plugin/log"
)

func NewConsumer(c []*config.RabbitmqConfig) *Consumer {
	return &Consumer{
		conf:       c,
		exitSignal: make(chan bool),
	}
}

type Consumer struct {
	conf        []*config.RabbitmqConfig
	mqCliMaps   sync.Map
	handlerMaps sync.Map // 原子map

	stopRun    bool
	exitSignal chan bool
	quantity   int32
}

func (cr *Consumer) Start() {
	for _, mqconf := range cr.conf {
		if _, ok := cr.handlerMaps.Load(mqconf.QueueName); !ok {
			panic("请先定义消费的入口函数handler，请检查！")
		}
	}

	go func() {
		// 每个配置开启对应的消费者
		for i, _ := range cr.conf {
			cr.runOneQueueConsumer(cr.conf[i])

		}
	}()
}

func (cr *Consumer) runOneQueueConsumer(c *config.RabbitmqConfig) {
	if 0 >= c.RunQuantity {
		return
	}

	param := &queue.InitParam{
		AmqpUrl:         c.AmqpUrl,
		QueueName:       c.QueueName,
		ExchangeName:    c.ExchangeName,
		ExchangeType:    c.ExchangeType,
		ExchangeDurable: c.ExchangeDurable,
		RoutingKey:      c.RoutingKey,
		Durable:         c.Durable,
	}

	for i := 0; i < c.RunQuantity; i++ {
		go cr.startQueueConsumerWhile(param)
	}
}

func (cr *Consumer) startQueueConsumerWhile(param *queue.InitParam) {
	if err := cr.checkParameter(param); nil != err {
		log.Error("启动参数配置错误导致消息队列无法启动，请检查！NewRabbitMqClient-err:", err)

		return
	}

	atomic.AddInt32(&cr.quantity, 1)

	for {
		if cr.stopRun {
			break
		}

		mq, err := cr.initQueue(param)
		if nil != err {
			time.Sleep(time.Second)
			continue
		}

		cr.mqCliMaps.Store(mq, true)

		if err = cr.startQueueConsumer(mq); nil != err {
			time.Sleep(time.Second)
		}

		cr.mqCliMaps.Delete(mq)
	}

	atomic.AddInt32(&cr.quantity, -1)
	cr.exitSignal <- true
}

func (cr *Consumer) checkParameter(param *queue.InitParam) error {
	if "" == param.AmqpUrl {
		return errors.New("AmqpUrl参数非法")
	}
	if "" == param.QueueName {
		return errors.New("QueueName参数非法")
	}
	if "" == param.ExchangeName {
		return errors.New("ExchangeName参数非法")
	}

	return nil
}

func (cr *Consumer) initQueue(param *queue.InitParam) (*queue.RabbitMqClient, error) {
	defer utils.CatchException("startQueueConsumer")

	mqCli, err := queue.NewRabbitMqClient(param)
	if nil != err {
		log.Error("消息队列无法启动，请检查！NewRabbitMqClient-err:", err)
		time.Sleep(3 * time.Second)

		return nil, err
	}

	return mqCli, nil
}

func (cr *Consumer) startQueueConsumer(mqCli *queue.RabbitMqClient) error {
	defer utils.CatchException("startQueueConsumer")
	defer mqCli.Close()

	v, _ := cr.handlerMaps.Load(mqCli.QueueName)
	consumerH := v.(*ConsumerStru)

	if err := mqCli.NewChannel(); nil != err {
		log.Error("消息队列无法启动，请检查！NewChannel-err:", err)
		time.Sleep(3 * time.Second)

		return err
	}

	if err := mqCli.Bind(); nil != err {
		log.Error("消息队列无法启动，请检查！Bind-err:", err)
		time.Sleep(3 * time.Second)

		return err
	}

	if err := mqCli.ConsumeSelfDefine(consumerH.QueueName, consumerH.AutoAck, consumerH.Handler); nil != err {
		log.Error("消息队列无法启动，请检查！Bind-err:", err)
		time.Sleep(3 * time.Second)

		return err
	}

	return nil
}

type ConsumerStru struct {
	QueueName string
	AutoAck   bool
	Handler   queue.ConsumerHandler
}

func (cr *Consumer) AddConsumerHandler(queueName string, autoAck bool, handler queue.ConsumerHandler) {
	cr.handlerMaps.Store(queueName, &ConsumerStru{
		QueueName: queueName,
		AutoAck:   autoAck,
		Handler:   handler,
	})
}

func (cr *Consumer) SecurityExit() {
	cr.stopRun = true

	// 发送退出信号
	cr.mqCliMaps.Range(func(key interface{}, value interface{}) bool {
		cli := key.(*queue.RabbitMqClient)

		cli.SecurityExit()

		return true
	})
	count := 0
	tm := time.Tick(3 * time.Second)
	for {

		select {
		case <-cr.exitSignal:
			if atomic.CompareAndSwapInt32(&cr.quantity, 0, 0) {
				return
			}
		case <-tm:
			count++
			if 3 < count {
				atomic.StoreInt32(&cr.quantity, 0)
				close(cr.exitSignal)
				return
			}
			log.Info("等待所有消费者退出。。。", count)

		}
	}

}
