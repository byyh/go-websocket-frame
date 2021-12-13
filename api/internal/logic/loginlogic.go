package logic

import (
	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"
)

type LoginLogic struct {
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &types.LoginResp{}, nil
}
