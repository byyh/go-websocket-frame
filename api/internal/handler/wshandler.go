package handler

import (
	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"
	"go-websocket-frame/api/ws"
	"go-websocket-frame/common/global/plugin/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WsHandler(ctx *gin.Context) {
	log.Info("WsHandler-begin.....")

	// todo : 检查token,用户信息指针放入ctx

	// test
	ctx.Set("user_info", &types.UserBase{
		Id:   10001,
		Name: "zhangsan",
	})

	// 调用websocket
	ws.ServeWs(ctx, ctx.Writer, ctx.Request)
}

func WsTestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// todo : 检查token,用户信息指针放入ctx

		// test
		http.ServeFile(w, r, "home22.html")
	}
}
