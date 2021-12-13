package handler

import (
	"go-websocket-frame/api/middleware"

	"github.com/gin-gonic/gin"
)

//
func InitRouter(router *gin.Engine) {

	// 测试
	// test := router.Group("/api/")
	// test.Use(middleware.LoginVerify)
	// {
	// 	test2 := test.Group("/v1.1/test")
	// 	{
	// 		test2.GET("/test", (&ctrAdminV1_1.Test{}).Test)
	// 		test2.GET("/test2", (&ctrAdminV1_1.Test{}).Test2)
	// 		test2.GET("/test3", (&ctrAdminV1_1.Test{}).Test3)
	// 		test2.GET("/test4", (&ctrAdminV1_1.Test{}).Test4)
	// 	}
	// 	exam := test.Group("/v1.1/exam")
	// 	{
	// 		exam.POST("/db", (&ctrAdminV1_1.Exam{}).Db)
	// 		exam.GET("/redis", (&ctrAdminV1_1.Exam{}).Redis)
	// 	}
	// }

	login := router.Group("/api/v1.1/login").Use(middleware.LoginVerify)
	{
		// login.GET("/verify", (&ctrAdminV1_1Login.Employee{}).Verify)
		// login.POST("/send/sms", (&ctrAdminV1_1Login.Employee{}).SendSms)
		// login.POST("/in", (&ctrAdminV1_1Login.Employee{}).In)
		// login.POST("/out", (&ctrAdminV1_1Login.Employee{}).Out)
		login.POST("/test", NewTest().Test)
		login.POST("/testdbi", NewTest().TestDbi)
		login.POST("/testdbs", NewTest().TestDbs)
		login.POST("/testdbu", NewTest().TestDbu)
		login.POST("/testrds", NewTest().TestRds)
	}

	ws := router.Group("/api/v1.1/ws")
	{
		ws.GET("/:token", WsHandler)
	}

}
