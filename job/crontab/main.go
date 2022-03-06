package main

import (
	"flag"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
	"go-websocket-frame/job/crontab/core"
	globalapi "go-websocket-frame/job/crontab/global"
	"go-websocket-frame/job/crontab/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/crontab.yaml", "the config file")

func main() {
	flag.Parse()

	var cf config.Config
	conf.MustLoad(*configFile, &cf)
	globalapi.InitConf(&cf)

	// 日志
	log.Init(cf.FluentLogConf)

	// init redis
	global.InitRedis(cf.RedisConf)

	// init mysql
	global.InitMysql(cf.MysqlDns)

	// 生产者
	global.InitProducer(cf.RabbitmqConf)

	core.NewProcess().Start()
}
