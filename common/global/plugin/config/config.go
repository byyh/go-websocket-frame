package config

type RedisConf struct {
	Host  string `json:"Host"`                                   // redis地址
	Type  string `json:"Type,default=node,options=node|cluster"` // redis类型
	Pass  string `json:"Pass,optional"`                          // redis密码
	Dbnum int    `json:"Dbnum"`
}

type RabbitmqConfig struct {
	AmqpUrl         string `json:"AmqpUrl"`
	QueueName       string `json:"QueueName"`
	ExchangeName    string `json:"ExchangeName"`
	ExchangeType    string `json:"ExchangeType"`
	ExchangeDurable bool   `json:"ExchangeDurable"`
	RoutingKey      string `json:"RoutingKey"`
	Durable         bool   `json:"Durable"`
	RunQuantity     int    `json:"RunQuantity"`
}

type FluentLogConfig struct {
	Isopen bool   `json:"Isopen"`
	Host   string `json:"Host"`
	Port   int    `json:"Port"`
	Tag    string `json:"Tag"`
}
