package svc

import (
	"context"
	"go-websocket-frame/api/internal/config"
)

type ServiceContext struct {
	Conf config.Config
	Ctx  context.Context
}

func NewServiceContext(c *config.Config, ctx context.Context) *ServiceContext {
	return &ServiceContext{
		Conf: *c,
		Ctx:  ctx,
	}
}
