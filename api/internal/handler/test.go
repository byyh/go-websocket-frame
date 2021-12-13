package handler

import (
	globalapi "go-websocket-frame/api/global"
	"go-websocket-frame/api/internal/logic"
	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"
	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/global/plugin/log"

	"github.com/gin-gonic/gin"
)

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}

func (h Test) Test(ctx *gin.Context) {
	log.Info("Test-begin.....")

	globalapi.Success(ctx, map[string]interface{}{
		"id":   10001,
		"name": "zhangsan",
	})
}

func (h Test) TestDbi(ctx *gin.Context) {
	log.Info("TestDbi-begin.....")

	id, err := logic.NewTestLogic(svc.NewServiceContext(globalapi.GetConf(), ctx)).
		TestDbi()
	if nil != err {
		globalapi.FailedByCode(ctx, err.Code())
	}

	globalapi.Success(ctx, map[string]interface{}{"id": id})
}

func (h Test) TestDbs(ctx *gin.Context) {
	log.Info("TestDbs-begin.....")

	id := ctx.GetInt64("id")

	ent, err := logic.NewTestLogic(svc.NewServiceContext(globalapi.GetConf(), ctx)).
		TestDbs(id)
	if nil != err {
		globalapi.FailedByCode(ctx, atom.ErrCodeDb)

		return
	}

	globalapi.Success(ctx, ent)
}

func (h Test) TestDbu(ctx *gin.Context) {
	log.Info("TestDbu-begin.....")

	var (
		req types.TestEmployeeUpdateReq
	)

	if err := ctx.ShouldBindJSON(&req); nil != err {
		globalapi.FailedByCode(ctx, atom.ErrCodeInput)

		return
	}

	err := logic.NewTestLogic(svc.NewServiceContext(globalapi.GetConf(), ctx)).
		TestDbu(&req)
	if nil != err {
		globalapi.FailedByCode(ctx, atom.ErrCodeDb)

		return
	}

	globalapi.Success(ctx, struct{}{})
}

func (h Test) TestRds(ctx *gin.Context) {
	log.Info("TestRds-begin.....")

	id := ctx.GetInt64("id")

	ent, err := logic.NewTestLogic(svc.NewServiceContext(globalapi.GetConf(), ctx)).
		TestRds(id)
	if nil != err {
		globalapi.FailedByCode(ctx, atom.ErrCodeDb)

		return
	}

	globalapi.Success(ctx, map[string]interface{}{
		"data": ent,
	})
}
