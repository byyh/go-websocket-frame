// 用户
package wslogic

import (
	"go-websocket-frame/api/internal/wssvc"
	"go-websocket-frame/common/atom"

	proto "go-websocket-frame/api/internal/wshandler/protopb"
)

type UserLogic struct {
	svcCtx *wssvc.ServiceContext
}

func NewUserLogic(svcCtx *wssvc.ServiceContext) *UserLogic {
	return &UserLogic{
		svcCtx: svcCtx,
	}
}

// 初始化本人信息
func (l UserLogic) InitSelf(req *proto.InitSelfReq, resp *proto.Empty) atom.Error {
	// todo: 该处添加逻辑代码

	// 如果发生错误，请返回 atom.Error，如下示例,错误编码根据时间情况传入
	// if ... {
	// 		return atom.NewMyErrorByCode(atom.ErrCodeCommon)
	// }

	return nil
}

// 获取本人信息
func (l UserLogic) GetSelfInfo(req *proto.Empty, resp *proto.UserInfoResp) atom.Error {
	// todo: 该处添加逻辑代码

	// 如果发生错误，请返回 atom.Error，如下示例,错误编码根据时间情况传入
	// if ... {
	// 		return atom.NewMyErrorByCode(atom.ErrCodeCommon)
	// }

	return nil
}

// 获取推荐用户列表
func (l UserLogic) GetRecommendUserList(req *proto.UserListReq, resp *proto.UserListResp) atom.Error {
	// todo: 该处添加逻辑代码

	// 如果发生错误，请返回 atom.Error，如下示例,错误编码根据时间情况传入
	// if ... {
	// 		return atom.NewMyErrorByCode(atom.ErrCodeCommon)
	// }

	return nil
}
