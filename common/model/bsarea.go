package model

type BsArea struct {
	Id    int64  `gorm:"primary_key;auto_increment;" json:"id"`
	Name  string `json:"name"`  // 地区名称
	Pid   string `json:"pid"`   // 上级id
	Sort  int8   `json:"sort"`  // 排序
	Level int8   `json:"level"` // 层级
}

func (BsArea) TableName() string {
	return "bs_area"
}
