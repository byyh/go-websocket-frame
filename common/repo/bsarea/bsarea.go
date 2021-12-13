
package bsarea

import (
	"fmt"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/model"
	"sync"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

var (
	r *repository
	m sync.RWMutex
)

type repository struct {
	db      *gorm.DB
	isCache bool
	rdsCli  *redis.Client
}

func GetRepository() (Repository, error) {
	if r == nil {
		return nil, fmt.Errorf("[GetRepository] 未初始化")
	}
	return r, nil
}

// Init 初始化仓储层
func init() {
	m.Lock()
	defer m.Unlock()

	if r != nil {
		return
	}

	r = &repository{
		db:      global.GetDb(),
		isCache: false,
		rdsCli:  global.GetRedis(),
	}
}

type Repository interface {
	// 根据 id 查询
	GetById(id int64) (*model.BsArea, error)

	//  根据 ids 查询列表
	GetListByIds(ids []int64) (res []model.BsArea, err error)

	// 创建
	Create(data *model.BsArea) (int64, error)

	// 编辑
	Update(data map[string]interface{}) error
	
	// barch编辑
	UpdateBatch(datas []map[string]interface{}) error

	// 删除
	Delete(ids []int64, updatedBy int64) error
	
	// 简单条件查询
	QueryByMaps(data map[string]interface{}, pageSize, current int, order string) (res []model.BsArea, err error)
	
	// 简单条件查询数量
	QueryTotalByMaps(data map[string]interface{}) (res int64, err error)
	
	// 简单条件查询
	QueryByWhere(pageSize, current int, order, where string, param ...interface{}) (res []model.BsArea, err error)
	
	// 简单条件查询数量
	QueryTotalByWhere(where string, param ...interface{}) (res int64, err error)
}

