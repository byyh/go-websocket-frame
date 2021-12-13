package logic

import (
	"fmt"
	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"
	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
	"go-websocket-frame/common/model"
	"go-websocket-frame/common/repo"
)

type TestLogic struct {
	svcCtx *svc.ServiceContext
}

func NewTestLogic(svcCtx *svc.ServiceContext) TestLogic {
	return TestLogic{
		svcCtx: svcCtx,
	}
}

func (h TestLogic) TestDbi() (int64, atom.Error) {
	log.Info("TestDbi-begin.....")

	m := model.BsShopEmployee{
		ShopId:   "110",
		Name:     "zs",
		Password: "dddddd",
		Mobile:   "13099998888",
	}

	id, err := repo.GetBsshopemployeeRepo().Create(&m)
	if nil != err {
		log.Error("GetBsshopemployeeRepo().Create-err", err)
		return 0, atom.NewMyErrorByCode(atom.ErrCodeDb)
	}

	return id, nil
}

func (h TestLogic) TestDbs(id int64) (*model.BsShopEmployee, atom.Error) {
	log.Info("TestDbs-begin.....")

	ent, err := repo.GetBsshopemployeeRepo().GetById(id)
	if nil != err {
		log.Error("GetBsshopemployeeRepo().GetById-err", err)
		return nil, atom.NewMyErrorByCode(atom.ErrCodeDb)
	}

	return ent, nil
}

func (h TestLogic) TestDbu(req *types.TestEmployeeUpdateReq) atom.Error {
	log.Info("TestDbu-begin.....")

	if err := repo.GetBsshopemployeeRepo().Update(map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
	}); nil != err {
		log.Error("GetBsshopemployeeRepo().Update-err", err)
		return atom.NewMyErrorByCode(atom.ErrCodeDb)
	}

	return nil
}

func (h TestLogic) TestRds(id int64) (string, atom.Error) {
	log.Info("TestRds-begin.....")

	rdsCli := global.GetRedis()
	key := fmt.Sprintf("%s-%s-%d", "social-db", model.BsShopEmployee{}.TableName(), id)
	res := rdsCli.Get(key)
	if err := res.Err(); nil != err {
		log.Error("TestLogic-rdsCli.Get-err", err)
		return "", atom.NewMyErrorByCode(atom.ErrCodeRedis)

	}

	return res.Val(), nil
}
