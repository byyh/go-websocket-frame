package global

import (
	"go-websocket-frame/common/atom"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, res interface{}) {
	ctx.JSON(200, gin.H{
		"code":    atom.Success,
		"message": atom.GetMsgByCode(atom.Success),
		"data":    res,
	})
}

func FailedByCode(ctx *gin.Context, code int) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": atom.GetMsgByCode(code),
		"data":    nil,
	})
}

// 返回失败
func Failed(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": msg,
		"data":    nil,
	})
}

// 返回非200HTTP状态码
func OutHttpError(ctx *gin.Context, httpCode int, msg string) {
	ctx.JSON(httpCode, gin.H{
		"message": msg,
	})
}
