// 定义api的全局初始化变量，各个业务层可直接使用

package global

import (
	"go-websocket-frame/api/internal/config"
)

func GetConf() *config.Config {
	return conf
}
