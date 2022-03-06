package bsarea

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
func (r *repository) Create(data *model.BsArea) (int64, error) {
	odb := r.db

	err := odb.Create(&data).Error
	if err != nil {
		log.Error("[repo.BsArea.Create] 创建发生错误，err：", err)
		return 0, errors.New("创建BsArea发生错误")
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

	if err := odb.Model(&model.BsArea{}).Where("id=?", data["id"]).
		Updates(data).Error; err != nil {
		log.Error("[repo.BsArea.Update] 编辑发生错误，err：", err)
		return errors.New("编辑BsArea发生错误")
	}

	if r.isCache {
		key := "social-db-BsArea"
		// 从缓存中获取数据
		lockKey := fmt.Sprintln("%s-%d", key, data["id"].(int64))

		// 锁定
		bl, err := cache.SetNxWait(lockKey, "ok", 3000)
		if nil != err {
			log.Error("BsArea.Update-cache.SetNxWait-err:", err)
			return err
		}
		if !bl {
			log.Error("BsArea.Update-cache.SetNxWait-failed:", bl)
			return errors.New("缓存锁定失败")
		}
		defer cache.UnLock(lockKey)

		// del cache
		if err := cache.HashDel(key, id); nil != err {
			log.Error("[repo.BsArea.Update-save-cache.HashDel] 删除缓存发生错误，err：", err)
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
		if err := tx.Model(&model.BsArea{}).Where("id=?", datas[i]["id"]).
			Updates(datas[i]).Error; err != nil {
			tx.Rollback()

			log.Error("[repo.BsArea.Update] 编辑发生错误，err：", err, data)
			return errors.New("batch编辑BsArea发生错误")
		}
	}

	tx.Commit()

	if r.isCache {
		key := "social-db-BsArea"

		for i, data := range datas {

			// 从缓存中获取数据
			lockKey := fmt.Sprintln("%s-%d", key, data["id"].(int64))

			// 锁定
			bl, err := cache.SetNxWait(lockKey, "ok", 3000)
			if nil != err {
				log.Error("BsArea.UpdateBatch-cache.SetNxWait-err:", err)
				return err
			}
			if !bl {
				log.Error("BsArea.UpdateBatch-cache.SetNxWait-failed:", bl)
				return errors.New("缓存锁定失败")
			}
			defer cache.UnLock(lockKey)

			// del cache
			if err := cache.HashDel(key, datas[i]["id"].(int64)); nil != err {
				log.Error("[repo.BsArea.UpdateBatch-save-cache.HashDel] 删除缓存发生错误，err：", err)
				// 对外不报错
				return nil
			}
		}
	}

	return nil
}

// 删除
func (r *repository) Delete(ids []int64, updatedBy int64) error {
	err := r.db.Model(model.BsArea{}).Where("id in (?)", ids).
		Updates(map[string]interface{}{
			"deleted_by": updatedBy,
			"deleted_at": time.Now(),
		}).Error
	if err != nil {
		log.Error("[repo.BsArea.Close] 关闭发生错误，err：", err)
		return errors.New("关闭BsArea发生错误")
	}

	if r.isCache {
		key := "social-db-BsArea"

		for _, id := range ids {

			// 从缓存中获取数据
			lockKey := fmt.Sprintln("%s-%d", key, id)

			// 锁定
			bl, err := cache.SetNxWait(lockKey, "ok", 3000)
			if nil != err {
				log.Error("BsArea.UpdateBatch-cache.SetNxWait-err:", err)
				return err
			}
			if !bl {
				log.Error("BsArea.UpdateBatch-cache.SetNxWait-failed:", bl)
				return errors.New("缓存锁定失败")
			}
			defer cache.UnLock(lockKey)

			// del cache
			if err := cache.HashDel(key, id); nil != err {
				log.Error("[repo.BsArea.UpdateBatch-save-cache.HashDel] 删除缓存发生错误，err：", err)
				// 对外不报错
				return nil
			}
		}
	}

	return nil
}

// 单个
func (r *repository) GetById(id int64) (*model.BsArea, error) {
	var data model.BsArea

	key := "social-db-BsArea"

	// 判断缓存
	if r.isCache {
		// 从缓存中获取数据
		lockKey := fmt.Sprintln("%s-%d", key, id)

		// 锁定
		bl, err := cache.SetNxWait(lockKey, "ok", 3000)
		if nil != err {
			log.Error("BsArea.GetById-cache.SetNxWait-err:", err)
			return nil, err
		}
		if !bl {
			log.Error("BsArea.GetById-cache.SetNxWait-failed:", bl)
			return nil, errors.New("缓存锁定失败")
		}
		defer cache.UnLock(lockKey)

		bl, err = cache.HExists("social-db-BsArea", id)
		if nil != err {
			log.Error("BsArea.GetById-cache.HExists-err:", err)
			return nil, err
		}
		if bl {
			res, err := cache.HashOneSearch(key, id)
			if nil != err {
				log.Error("BsArea.GetById-cache.HashOneSearch-err:", err)
				return nil, err
			}

			if err := json.Unmarshal([]byte(res), &data); nil != err {
				log.Error("BsArea.GetById-json.Unmarshal-err:", err)
				return nil, err
			}

			return &data, nil
		}
	}

	err := r.db.Model(&model.BsArea{}).Where("id=?", id).First(&data).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errors.New("数据不存在") // 不存在情况
		}
		log.Error("[repo.BsArea.GetById] 查询指定活动发生错误，err：", err)
		return nil, err // 包含不存在情况
	}

	if r.isCache {
		bt, err := json.Marshal(&data)
		if nil != err {
			log.Error("[repo.BsArea.GetById-save-json.Marshal] 保存缓存发生错误，err：", err)
			// 对外不报错
			return &data, nil
		}

		// save cache
		if err := cache.HashAddTtl(key, id, string(bt), atom.DbCacheExpireTms); nil != err {
			log.Error("[repo.BsArea.GetById-save-cache.HashAddTtl] 保存缓存发生错误，err：", err)
			// 对外不报错
			return &data, nil
		}
	}

	return &data, nil
}

// 列表
func (r *repository) GetListByIds(ids []int64) (res []model.BsArea, err error) {
	err = r.db.Model(&model.BsArea{}).Where("id in (?)", ids).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsArea.GetListByIds] 查询list发生错误:", err)
		return
	}

	return
}

func (r *repository) QueryByMaps(data map[string]interface{},
	pageSize, current int, order string) (res []model.BsArea, err error) {
	offset := (current - 1) * pageSize
	err = r.db.Model(&model.BsArea{}).
		Where(data).Order(order).
		Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsArea.QueryByMaps] 查询list发生错误:", err)
	}

	return
}

func (r *repository) QueryTotalByMaps(data map[string]interface{}) (res int64, err error) {
	err = r.db.Model(&model.BsArea{}).
		Where(data).
		Count(&res).Error
	if err != nil {
		log.Error("[repo.BsArea.QueryTotalByMaps] 查询total发生错误:", err)
	}

	return
}

func (r *repository) QueryByWhere(pageSize, current int, order, where string,
	param ...interface{}) (res []model.BsArea, err error) {
	offset := (current - 1) * pageSize

	err = r.db.Model(&model.BsArea{}).
		Where(where, param...).Order(order).
		Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		log.Error("[repo.BsArea.QueryByWhere] 查询total发生错误:", err)
	}

	return
}

func (r *repository) QueryTotalByWhere(where string,
	param ...interface{}) (res int64, err error) {
	err = r.db.Model(&model.BsArea{}).
		Where(where, param...).
		Count(&res).Error
	if err != nil {
		log.Error("[repo.BsArea.QueryTotalByWhere] 查询total发生错误:", err)
	}

	return
}
