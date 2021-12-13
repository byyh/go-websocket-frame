package logic

import (
	"context"

	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostTestLogic {
	return PostTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() error {
	// todo: add your logic here and delete this line

	return nil
}
