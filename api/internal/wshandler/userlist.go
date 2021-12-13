// 用户列表
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

// 获取推荐用户列表
func UserListGetRecommendList(c *ws.Client, cmd uint32, reqUid int64, data []byte) atom.Error {
	var (
		req  proto.RecommendReq
		resp proto.UserListResp
	)

	if err := sproto.Unmarshal(data, &req); nil != err {
		log.Errorf("UserListGetRecommendList-Unmarshal-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, string(data), err)
		return atom.NewMyErrorByCode(atom.ErrCodeJson)
	}

	// call logic
	conf := global.GetConf()
	user := c.GetUser()
	svcCtx := wssvc.NewServiceContext(conf, user)

	l := wslogic.NewUserListLogic(svcCtx)
	err := l.GetRecommendList(&req, &resp)
	if err != nil {
		log.Error("UserList-logic.GetRecommendList-err,cmd=%d,reqUid=%d,c.user=%s,data=%s,err=%s",
			cmd, reqUid, c.User, req.String(), err)
		c.SendErrorMsg(cmd, reqUid, err.Code())

		return nil
	}

	// ret cli
	c.SendMsg(cmd, reqUid, &resp)

	return nil
}
