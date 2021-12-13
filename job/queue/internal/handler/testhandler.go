package handler

import (
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

// 测试样例
func TestHandler1(data []byte) error {
	logx.Info("begin TestHandler1.run:", string(data))

	// l := logic.NewTestLogic(r.Context(), h.Ctx)
	// resp, err := l.Test()
	//time.Sleep(1 * time.Millisecond)

	logx.Info("end TestHandler1.end")

	return nil
}

func TestHandler2(data []byte) error {
	logx.Info("begin TestHandler2.run:", string(data))

	// l := logic.NewTestLogic(r.Context(), h.Ctx)
	// resp, err := l.Test()
	time.Sleep(1 * time.Millisecond)

	logx.Info("end TestHandler2.end")

	return nil
}

func TestHandler3(data []byte) error {
	logx.Info("begin TestHandler3.run:", string(data))

	// l := logic.NewTestLogic(r.Context(), h.Ctx)
	// resp, err := l.Test()
	time.Sleep(1 * time.Millisecond)

	logx.Info("end TestHandler3.end")

	return nil
}
