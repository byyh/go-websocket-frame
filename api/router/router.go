package router

import (
	"go-websocket-frame/api/global"
	"go-websocket-frame/api/internal/handler"
	"go-websocket-frame/api/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func initBase() {
	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Recover) // 自定义全局的异常检测

	cf := global.GetConf()
	if "debug" == cf.RunMode || "test" == cf.RunMode {
		pprof.Register(router) // 性能监控，用于开发测试环境
	}
}

func Init() *gin.Engine {
	initBase()

	// 初始化路由
	handler.InitRouter(router)

	return router
}
