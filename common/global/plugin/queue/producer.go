package queue

import (
	"go-websocket-frame/common/global/plugin/config"
	"go-websocket-frame/common/plugin/queue"
)

type ProducerIface interface {
	Init() error
	GetMq() queue.RabbitmqIface
}

type Producer struct {
	conf  *config.RabbitmqConfig
	MqCli *queue.RabbitMqClient
}

func NewProducer(c *config.RabbitmqConfig) *Producer {
	return &Producer{
		conf: c,
	}
}

func (p *Producer) Init() error {
	var err error

	c := p.conf
	p.MqCli, err = queue.NewRabbitMqClient(&queue.InitParam{
		AmqpUrl:         c.AmqpUrl,
		QueueName:       c.QueueName,
		ExchangeName:    c.ExchangeName,
		ExchangeType:    c.ExchangeType,
		ExchangeDurable: c.ExchangeDurable,
		RoutingKey:      c.RoutingKey,
		Durable:         c.Durable,
	})
	if nil != err {
		return err
	}

	return nil
}

func (p *Producer) GetMq() queue.RabbitmqIface {
	return p.MqCli
}
