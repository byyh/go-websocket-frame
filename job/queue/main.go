package main

import (
	"flag"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
	globalapi "go-websocket-frame/job/queue/global"
	"go-websocket-frame/job/queue/internal/config"
	"go-websocket-frame/job/queue/internal/handler"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/queue.yaml", "the config file")

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

	// 启动消费
	cr := global.InitConsumer(cf.RabbitmqConfList)
	handler.InitRoute(cr)
	cr.Start()

	// 生产者
	global.InitProducer(cf.RabbitmqConfList[0])

	time.Sleep(1 * time.Second)

	//go handler.TestProducter()

	// 捕获退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("正在关闭...")
	// 优雅关闭
	cr.SecurityExit()

	log.Info("结束.")

}
