// 用户列表
package wslogic

import (
	"go-websocket-frame/api/internal/wssvc"
	"go-websocket-frame/common/atom"

	proto "go-websocket-frame/api/internal/wshandler/protopb"
)

type UserListLogic struct {
	svcCtx *wssvc.ServiceContext
}

func NewUserListLogic(svcCtx *wssvc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		svcCtx: svcCtx,
	}
}

// 获取推荐用户列表
func (l UserListLogic) GetRecommendList(req *proto.RecommendReq, resp *proto.UserListResp) atom.Error {
	// todo: 该处添加逻辑代码

	// 如果发生错误，请返回 atom.Error，如下示例,错误编码根据时间情况传入
	// if ... {
	// 		return atom.NewMyErrorByCode(atom.ErrCodeCommon)
	// }

	return nil
}
