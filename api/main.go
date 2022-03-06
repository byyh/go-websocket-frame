package main

import (
	"context"
	"flag"
	globalapi "go-websocket-frame/api/global"
	"go-websocket-frame/api/internal/config"
	"go-websocket-frame/api/internal/consumer"
	"go-websocket-frame/api/internal/wshandler"
	"go-websocket-frame/api/ws"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"go-websocket-frame/api/router"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/social-api.yaml", "the config file")

func main() {
	flag.Parse()

	var cf config.Config
	conf.MustLoad(*configFile, &cf)
	globalapi.InitConf(&cf)
	global.InitLog(cf.FluentLogConf)

	gin.SetMode(cf.RunMode)
	router := router.Init()

	//router.Use(gin.Logger())

	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(cf.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 14, // 16k左右
	}

	// 捕获退出信号
	quit := make(chan os.Signal)

	// init redis
	global.InitRedis(cf.RedisConf)

	// init mysql
	global.InitMysql(cf.MysqlDns)

	// init 消费
	cr := global.InitConsumer(cf.RabbitmqConfList)
	consumer.InitRoute(cr)
	cr.Start()

	// init 生产者
	global.InitProducer(cf.RabbitmqConfList[0])

	// 启动ws通道协程
	go func() {
		// ws路由规则
		wshandler.InitRoute()
		ws.NewHub().Run()

	}()

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
			signal.Stop(quit)
		}
	}()

	go func() {
		tms := time.Tick(5 * time.Second)
		for {
			select {
			case <-tms:
				log.Info("当前协程数量：", runtime.NumGoroutine())
			}
		}
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("正在关闭服务...")
	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("服务关闭异常：", err)
	}
	log.Info("服务结束.")
}
