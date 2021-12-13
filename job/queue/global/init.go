package global

import (
	"go-websocket-frame/job/queue/internal/config"
)

var (
	conf *config.Config
)

func InitConf(c *config.Config) {
	if nil == c {
		panic("InitConf配置参数不能为空")
	}

	conf = c
}
