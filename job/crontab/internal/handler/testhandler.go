package wshandler

import (
	"go-websocket-frame/job/crontab/core/iface"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestHandler struct {
	iface.ProcIface // 一定要实现该接口
	Ctx             *svc.ServiceContext
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h TestHandler) Init(ctx *svc.ServiceContext) { h.Ctx = ctx }
func (h TestHandler) Destructor()                  {}

// 测试样例
func (h TestHandler) Run() {
	logx.Info("begin TestHandler.run")

	// l := logic.NewTestLogic(r.Context(), h.Ctx)
	// resp, err := l.Test()
	time.Sleep(10 * time.Second)

	logx.Info("end TestHandler.run")

	return nil
}
