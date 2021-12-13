package middleware

// 捕获全局异常

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			// 系统级别的异常，打印错误栈
			debug.PrintStack()

			c.JSON(http.StatusBadRequest, gin.H{
				"code": "-10",
				"msg":  errorToString(err),
				"data": nil,
			})

			//终止后续接口调用
			c.Abort()
		}
	}()

	// 后继
	c.Next()
}

// recover错误信息转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
