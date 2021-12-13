package utils

import (
	"runtime/debug"

	"go-websocket-frame/common/global/plugin/log"
)

func CatchException(title string) {
	if err := recover(); nil != err {
		// 写错误异常日志
		log.Error("exec ", title, " error: handle recover", err)
		log.Error("异常栈：", string(debug.Stack()))
	}
}

func CatchExceptionFunc(title string, ef ExceptionFunc) {
	if err := recover(); nil != err {
		// 写错误异常日志
		log.Error("exec ", title, " error: handle recover", err)
		log.Error("异常栈：", string(debug.Stack()))

		ef()
	}
}

type ExceptionFunc func()
