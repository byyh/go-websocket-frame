package handler

import (
	"go-websocket-frame/job/crontab/core"
)

func InitRoute() {
	// 唯一标记     运行规则：秒 分 时 日 月 周（0代表周日）
	core.AddHandleTask("TestHandler", "*/50 * * * * *", core.NewRunProcess(TestHandler)) // demo
}
