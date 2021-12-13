package types

// 用户信息
type UserBase struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	// todo 其他信息
}
