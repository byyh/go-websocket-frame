// 用户
package wshandler

import (
	"go-websocket-frame/api/global"
	proto "go-websocket-frame/api/internal/wshandler/protopb"
	"go-websocket-frame/api/internal/wslogic"
	"go-websocket-frame/api/internal/wssvc"
	"go-websocket-frame/api/ws"
	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/global/plugin/log"

	sproto "github.com/golang/protobuf/proto"
)

// 初始化本人信息
func UserInitSelf(c *ws.Client, cmd uint32, reqUid int64, data []byte) atom.Error {
	var (
		req  proto.InitSelfReq
		resp proto.Empty
	)

	if err := sproto.Unmarshal(data, &req); nil != err {
		log.Errorf("UserInitSelf-Unmarshal-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, string(data), err)
		return atom.NewMyErrorByCode(atom.ErrCodeJson)
	}

	// call logic
	conf := global.GetConf()
	user := c.GetUser()
	svcCtx := wssvc.NewServiceContext(conf, user)

	l := wslogic.NewUserLogic(svcCtx)
	err := l.InitSelf(&req, &resp)
	if err != nil {
		log.Error("User-logic.InitSelf-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, req.String(), err)
		c.SendErrorMsg(cmd, reqUid, err.Code())

		return nil
	}

	// ret cli
	c.SendMsg(cmd, reqUid, &resp)

	return nil
}

// 获取本人信息
func UserGetSelfInfo(c *ws.Client, cmd uint32, reqUid int64, data []byte) atom.Error {
	var (
		req  proto.Empty
		resp proto.UserInfoResp
	)

	if err := sproto.Unmarshal(data, &req); nil != err {
		log.Errorf("UserGetSelfInfo-Unmarshal-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, string(data), err)
		return atom.NewMyErrorByCode(atom.ErrCodeJson)
	}

	// call logic
	conf := global.GetConf()
	user := c.GetUser()
	svcCtx := wssvc.NewServiceContext(conf, user)

	l := wslogic.NewUserLogic(svcCtx)
	err := l.GetSelfInfo(&req, &resp)
	if err != nil {
		log.Error("User-logic.GetSelfInfo-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, req.String(), err)
		c.SendErrorMsg(cmd, reqUid, err.Code())

		return nil
	}

	// ret cli
	c.SendMsg(cmd, reqUid, &resp)

	return nil
}

// 获取推荐用户列表
func UserGetRecommendUserList(c *ws.Client, cmd uint32, reqUid int64, data []byte) atom.Error {
	var (
		req  proto.UserListReq
		resp proto.UserListResp
	)

	if err := sproto.Unmarshal(data, &req); nil != err {
		log.Errorf("UserGetRecommendUserList-Unmarshal-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, string(data), err)
		return atom.NewMyErrorByCode(atom.ErrCodeJson)
	}

	// call logic
	conf := global.GetConf()
	user := c.GetUser()
	svcCtx := wssvc.NewServiceContext(conf, user)

	l := wslogic.NewUserLogic(svcCtx)
	err := l.GetRecommendUserList(&req, &resp)
	if err != nil {
		log.Error("User-logic.GetRecommendUserList-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, req.String(), err)
		c.SendErrorMsg(cmd, reqUid, err.Code())

		return nil
	}

	// ret cli
	c.SendMsg(cmd, reqUid, &resp)

	return nil
}
