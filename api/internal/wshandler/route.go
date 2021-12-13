// ws router
package wshandler

import (
	proto "go-websocket-frame/api/internal/wshandler/protopb"
	"go-websocket-frame/api/ws"
)

func InitRoute() {
	// 所有的ws处理路由添加到这里

	// 用户列表 获取推荐用户列表
	ws.AddWsHandleFunc((uint32(proto.CmdBase_CmdBaseUserList))|uint32(proto.CmdUserList_CmdUserListGetRecommendList), UserListGetRecommendList)

	// 用户 初始化本人信息
	ws.AddWsHandleFunc((uint32(proto.CmdBase_CmdBaseUser))|uint32(proto.CmdUser_CmdUserInitSelf), UserInitSelf)

	// 用户 获取本人信息
	ws.AddWsHandleFunc((uint32(proto.CmdBase_CmdBaseUser))|uint32(proto.CmdUser_CmdUserGetSelfInfo), UserGetSelfInfo)

	// 用户 获取推荐用户列表
	ws.AddWsHandleFunc((uint32(proto.CmdBase_CmdBaseUser))|uint32(proto.CmdUser_CmdUserGetRecommendUserList), UserGetRecommendUserList)

}
