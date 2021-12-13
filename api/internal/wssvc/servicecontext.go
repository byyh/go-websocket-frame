package wssvc

import (
	"go-websocket-frame/api/internal/config"
	"go-websocket-frame/api/internal/types"
)

type ServiceContext struct {
	Conf config.Config
	User types.UserBase
}

func NewServiceContext(c *config.Config, user *types.UserBase) *ServiceContext {
	return &ServiceContext{
		Conf: *c,
		User: *user,
	}
}
