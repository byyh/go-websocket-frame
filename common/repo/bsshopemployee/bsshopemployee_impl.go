package bsshopemployee

import (
	"errors"
	"fmt"
	"time"

	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/global/plugin/log"
	"go-websocket-frame/common/model"
	"go-websocket-frame/common/utils/cache"

	json "github.com/json-iterator/go"
	"gorm.io/gorm"
)

// 创建
func (r *repository) Create(data *model.BsShopEmployee) (int64, error) {
	odb := r.db

	err := odb.Create(&data).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.Create] 创建发生错误，err：", err)
		return 0, errors.New("创建BsShopEmployee发生错误")
	}

	return data.Id, nil
}

// 编辑
func (r *repository) Update(data map[string]interface{}) error {
	id, ok := data["id"].(int64)
	if !ok {
		return errors.New("必须上传id参数")
	}

	odb := r.db

	if err := odb.Model(&model.BsShopEmployee{}).Where("id=?", data["id"]).
		Updates(data).Error; err != nil {
		log.Error("[repo.BsShopEmployee.Update] 编辑发生错误，err：", err)
		return errors.New("编辑BsShopEmployee发生错误")
	}

	if r.isCache {
		key := "social-db-BsShopEmployee"
		// 从缓存中获取数据
		lockKey := fmt.Sprintln("%s-%d", key, data["id"].(int64))

		// 锁定
		bl, err := cache.SetNxWait(lockKey, "ok", 3000)
		if nil != err {
			log.Error("BsShopEmployee.Update-cache.SetNxWait-err:", err)
			return err
		}
		if !bl {
			log.Error("BsShopEmployee.Update-cache.SetNxWait-failed:", bl)
			return errors.New("缓存锁定失败")
		}
		defer cache.UnLock(lockKey)

		// del cache
		if err := cache.HashDel(key, id); nil != err {
			log.Error("[repo.BsShopEmployee.Update-save-cache.HashDel] 删除缓存发生错误，err：", err)
			// 对外不报错
			return nil
		}
	}

	return nil
}

// barch编辑
func (r *repository) UpdateBatch(datas []map[string]interface{}) error {
	for _, data := range datas {
		if _, ok := data["id"]; !ok {
			return errors.New("必须上传id参数")
		}
	}

	tx := r.db.Begin()

	for i, data := range datas {
		if err := tx.Model(&model.BsShopEmployee{}).Where("id=?", datas[i]["id"]).
			Updates(datas[i]).Error; err != nil {
			tx.Rollback()

			log.Error("[repo.BsShopEmployee.Update] 编辑发生错误，err：", err, data)
			return errors.New("batch编辑BsShopEmployee发生错误")
		}
	}

	tx.Commit()

	if r.isCache {
		key := "social-db-BsShopEmployee"

		for i, data := range datas {

			// 从缓存中获取数据
			lockKey := fmt.Sprintln("%s-%d", key, data["id"].(int64))

			// 锁定
			bl, err := cache.SetNxWait(lockKey, "ok", 3000)
			if nil != err {
				log.Error("BsShopEmployee.UpdateBatch-cache.SetNxWait-err:", err)
				return err
			}
			if !bl {
				log.Error("BsShopEmployee.UpdateBatch-cache.SetNxWait-failed:", bl)
				return errors.New("缓存锁定失败")
			}
			defer cache.UnLock(lockKey)

			// del cache
			if err := cache.HashDel(key, datas[i]["id"].(int64)); nil != err {
				log.Error("[repo.BsShopEmployee.UpdateBatch-save-cache.HashDel] 删除缓存发生错误，err：", err)
				// 对外不报错
				return nil
			}
		}
	}

	return nil
}

// 删除
func (r *repository) Delete(ids []int64, updatedBy int64) error {
	err := r.db.Model(model.BsShopEmployee{}).Where("id in (?)", ids).
		Updates(map[string]interface{}{
			"deleted_by": updatedBy,
			"deleted_at": time.Now(),
		}).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.Close] 关闭发生错误，err：", err)
		return errors.New("关闭BsShopEmployee发生错误")
	}

	if r.isCache {
		key := "social-db-BsShopEmployee"

		for _, id := range ids {

			// 从缓存中获取数据
			lockKey := fmt.Sprintln("%s-%d", key, id)

			// 锁定
			bl, err := cache.SetNxWait(lockKey, "ok", 3000)
			if nil != err {
				log.Error("BsShopEmployee.UpdateBatch-cache.SetNxWait-err:", err)
				return err
			}
			if !bl {
				log.Error("BsShopEmployee.UpdateBatch-cache.SetNxWait-failed:", bl)
				return errors.New("缓存锁定失败")
			}
			defer cache.UnLock(lockKey)

			// del cache
			if err := cache.HashDel(key, id); nil != err {
				log.Error("[repo.BsShopEmployee.UpdateBatch-save-cache.HashDel] 删除缓存发生错误，err：", err)
				// 对外不报错
				return nil
			}
		}
	}

	return nil
}

// 单个
func (r *repository) GetById(id int64) (*model.BsShopEmployee, error) {
	var data model.BsShopEmployee

	key := "social-db-BsShopEmployee"

	// 判断缓存
	if r.isCache {
		// 从缓存中获取数据
		lockKey := fmt.Sprintln("%s-%d", key, id)

		// 锁定
		bl, err := cache.SetNxWait(lockKey, "ok", 3000)
		if nil != err {
			log.Error("BsShopEmployee.GetById-cache.SetNxWait-err:", err)
			return nil, err
		}
		if !bl {
			log.Error("BsShopEmployee.GetById-cache.SetNxWait-failed:", bl)
			return nil, errors.New("缓存锁定失败")
		}
		defer cache.UnLock(lockKey)

		bl, err = cache.HExists("social-db-BsShopEmployee", id)
		if nil != err {
			log.Error("BsShopEmployee.GetById-cache.HExists-err:", err)
			return nil, err
		}
		if bl {
			res, err := cache.HashOneSearch(key, id)
			if nil != err {
				log.Error("BsShopEmployee.GetById-cache.HashOneSearch-err:", err)
				return nil, err
			}

			if err := json.Unmarshal([]byte(res), &data); nil != err {
				log.Error("BsShopEmployee.GetById-json.Unmarshal-err:", err)
				return nil, err
			}

			return &data, nil
		}
	}

	err := r.db.Model(&model.BsShopEmployee{}).Where("id=?", id).First(&data).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errors.New("数据不存在") // 不存在情况
		}
		log.Error("[repo.BsShopEmployee.GetById] 查询指定活动发生错误，err：", err)
		return nil, err // 包含不存在情况
	}

	if r.isCache {
		bt, err := json.Marshal(&data)
		if nil != err {
			log.Error("[repo.BsShopEmployee.GetById-save-json.Marshal] 保存缓存发生错误，err：", err)
			// 对外不报错
			return &data, nil
		}

		// save cache
		if err := cache.HashAddTtl(key, id, string(bt), atom.DbCacheExpireTms); nil != err {
			log.Error("[repo.BsShopEmployee.GetById-save-cache.HashAddTtl] 保存缓存发生错误，err：", err)
			// 对外不报错
			return &data, nil
		}
	}

	return &data, nil
}

// 列表
func (r *repository) GetListByIds(ids []int64) (res []model.BsShopEmployee, err error) {
	err = r.db.Model(&model.BsShopEmployee{}).Where("id in (?)", ids).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.GetListByIds] 查询list发生错误:", err)
		return
	}

	return
}

func (r *repository) QueryByMaps(data map[string]interface{},
	pageSize, current int, order string) (res []model.BsShopEmployee, err error) {
	offset := (current - 1) * pageSize
	err = r.db.Model(&model.BsShopEmployee{}).
		Where(data).Order(order).
		Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.QueryByMaps] 查询list发生错误:", err)
	}

	return
}

func (r *repository) QueryTotalByMaps(data map[string]interface{}) (res int64, err error) {
	err = r.db.Model(&model.BsShopEmployee{}).
		Where(data).
		Count(&res).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.QueryTotalByMaps] 查询total发生错误:", err)
	}

	return
}

func (r *repository) QueryByWhere(pageSize, current int, order, where string,
	param ...interface{}) (res []model.BsShopEmployee, err error) {
	offset := (current - 1) * pageSize

	err = r.db.Model(&model.BsShopEmployee{}).
		Where(where, param...).Order(order).
		Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.QueryByWhere] 查询total发生错误:", err)
	}

	return
}

func (r *repository) QueryTotalByWhere(where string,
	param ...interface{}) (res int64, err error) {
	err = r.db.Model(&model.BsShopEmployee{}).
		Where(where, param...).
		Count(&res).Error
	if err != nil {
		log.Error("[repo.BsShopEmployee.QueryTotalByWhere] 查询total发生错误:", err)
	}

	return
}
