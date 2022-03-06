package model

import (
	"time"
)

type BsShopEmployee struct {
	Id           int64      `gorm:"primary_key;auto_increment;" json:"id"`
	ShopId       string     `json:"shop_id"`       // 商铺id，刚注册时没有，默认为 not-init
	LoginAccount string     `json:"login_account"` // 登陆账户，全网唯一
	Name         string     `json:"name"`          // 员工姓名
	Password     string     `json:"password"`      // 登陆密码
	InitPwd      string     `json:"init_pwd"`      // 初始密码，明文，初始化和重置使用
	Mobile       string     `json:"mobile"`        // 手机
	Code         string     `json:"code"`          // 员工编号
	Status       int8       `json:"status"`        // 状态，1有效，0初始化，-1无效
	IsDel        int8       `json:"is_del"`        // 是否删除
	Creator      string     `json:"creator"`       // 创建人
	Updater      string     `json:"updater"`       // 修改人
	CreateTime   *time.Time `json:"create_time"`   // 创建时间
	UpdateTime   *time.Time `json:"update_time"`   // 修改时间
}

func (BsShopEmployee) TableName() string {
	return "bs_shop_employee"
}
