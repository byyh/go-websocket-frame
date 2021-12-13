package cache

import (
	"errors"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
	"strconv"
	"time"

	"github.com/byyh/go/com"
	"github.com/go-redis/redis/v7"
)

const (
	HashFieldPrefix         = ""
	RedisFailedRepeatWaitMs = 50
	MaxPageSize             = 50
)

var (
	rdsCli = global.GetRedis()
)

//加锁
func SetNx(key, value string, ttl int64) (bl bool, err error) {
	//设置过期时间
	res := rdsCli.SetNX(key, value, time.Duration(ttl*1e9))
	if err = res.Err(); err != nil {
		return
	}

	return res.Val(), nil
}

//解锁
func UnLock(key string) (bool, error) {
	res := rdsCli.Del(key)
	if err := res.Err(); err != nil {
		return false, err
	}
	return true, nil
}

//
func SetNxWait(key, value string, ttl int64) (bl bool, err error) {
	for i := 0; i < 3; i++ {
		//设置过期时间
		res := rdsCli.SetNX(key, value, time.Duration(ttl)*1e6)
		if err = res.Err(); err != nil {
			time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
			continue
		}

		if bl = res.Val(); !bl {
			time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
			continue
		} else {
			return true, nil
		}
	}

	return

}

func IsRedisLock(key string) (bool, error) {
	_, err := rdsCli.Get(key).Result()
	if err == redis.Nil {
		//fmt.Println("key2 does not exist")
		return false, nil
	} else if err != nil {
		// 错误
		log.Error("[CooCalanderAdd.redis]错误", err)
		return false, err
	} else {
		return true, nil
	}
}

// 设置秒级过期时间
func Expire(key string, tms int64) error {
	return rdsCli.Expire(key, time.Duration(tms)*1e9).Err()
}

func Exists(key string) (bool, error) {
	res, err := rdsCli.Exists(key).Result()
	if nil != err {
		return false, err
	}
	if 0 == res {
		return false, nil
	}

	return true, nil
}

func HExists(key string, id int64) (bool, error) {
	field := HashFieldPrefix + strconv.FormatInt(id, 10)

	bl, err := rdsCli.HExists(key, field).Result()
	if nil != err {
		return false, err
	}

	return bl, nil
}

// tms 秒
func HashAddTtl(key string, id int64, jsonStr string, tms int64) (err error) {
	bl, err := Exists(key)
	if nil != err {
		return err
	}

	if err := HashAdd(key, id, jsonStr); nil != err {
		return err
	}

	if !bl {
		if err := Expire(key, tms); nil != err {
			log.Error("redis Expire err:", err)
		}
	}

	return nil
}

func HashAdd(key string, id int64, jsonStr string) (err error) {
	field := HashFieldPrefix + strconv.FormatInt(id, 10)

	for i := 0; i < 3; i++ {
		err = rdsCli.HSet(key, field, jsonStr).Err()
		if nil == err {
			return nil
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("保存list数据到redis失败：", err, key, jsonStr)
	}

	return
}

func HashDel(key string, id int64) (err error) {
	field := HashFieldPrefix + strconv.FormatInt(id, 10)

	for i := 0; i < 3; i++ {
		err = rdsCli.HDel(key, field).Err()
		if nil == err {
			return nil
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("从redis删除list数据失败：", err, key, id)
	}

	return
}

func HashClear(key string) (err error) {
	res := rdsCli.HKeys(key)
	if nil != res.Err() {
		log.Error("redis HashClear失败：", key, res.Err())
		return res.Err()
	}

	for _, field := range res.Val() {
		for i := 0; i < 3; i++ {
			err = rdsCli.HDel(key, field).Err()
			if nil == err {
				break
			}
			time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
		}

		if nil != err {
			log.Error("从redis清空hash数据失败：", err, key)

			return err
		}
	}

	return
}

func HashListSearch(key string, ids []int64) (res []string, err error) {
	var fields []string
	for _, v := range ids {
		fields = append(fields, HashFieldPrefix+strconv.FormatInt(v, 10))
	}

	rdsRes := rdsCli.HMGet(key, fields...)
	if nil != rdsRes.Err() {

		return res, rdsRes.Err()
	}
	log.Debug("fields===", fields)
	for _, v := range rdsRes.Val() {
		if nil == v {
			continue
		}
		if "string" != com.Typeof(v) {
			log.Error("redis return data type error:", v)
			err = errors.New("redis return data type error")
			return
		}

		res = append(res, v.(string))
	}

	return
}

func HashOneSearch(key string, id int64) (res string, err error) {
	field := HashFieldPrefix + strconv.FormatInt(id, 10)

	rdsRes := rdsCli.HGet(key, field)
	if redis.Nil == rdsRes.Err() {
		return "", nil
	} else if nil != rdsRes.Err() {
		return res, rdsRes.Err()
	}

	return rdsRes.Val(), nil
}

func HashAllSearch(key string, ids []int64) (res map[string]string, err error) {
	rdsRes := rdsCli.HGetAll(key)
	if nil != rdsRes.Err() {
		return res, rdsRes.Err()
	}

	return rdsRes.Val(), nil
}

func ListAdd(key string, val interface{}) (err error) {
	var (
		score float64
	)
	if "int64" == com.Typeof(val) {
		score = float64(val.(int64))
	}

	for i := 0; i < 3; i++ {
		err = rdsCli.ZAdd(key, &redis.Z{
			Score:  score,
			Member: val,
		}).Err()
		if nil == err {
			return nil
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("保存list数据到redis失败：", err, key)
	}

	return
}

func ListDel(key string, val interface{}) (err error) {
	for i := 0; i < 3; i++ {
		err = rdsCli.ZRem(key, val).Err()
		if nil == err {
			return nil
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("从redis删除list数据失败：", err, key)
	}

	return
}

func ListClear(key string) (err error) {
	res := rdsCli.ZCard(key)
	if nil != res.Err() {
		log.Error("redis ListClear error：", res.Err(), key)
		return res.Err()
	}

	if 0 >= res.Val() {
		return
	}

	for i := 0; i < 3; i++ {
		err = rdsCli.ZRemRangeByRank(key, 0, res.Val()).Err()
		if nil == err {
			return
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("从redis清空zsort数据失败：", err, key)

		return err
	}

	return
}

func ListSearch(key string, start, end int64, isDesc bool) (res []string, err error) {
	var (
		rdsRes *redis.StringSliceCmd
	)

	if MaxPageSize < (end - start) {
		end = start + MaxPageSize
	}

	if isDesc {
		rdsRes = rdsCli.ZRevRange(key, start, end)
	} else {
		rdsRes = rdsCli.ZRange(key, start, end)
	}

	if nil != rdsRes.Err() {
		return res, rdsRes.Err()
	}

	return rdsRes.Val(), nil
}

func ListTotal(key string) (res int64, err error) {
	rdsRe := rdsCli.ZCard(key)
	if nil != rdsRe.Err() {
		return res, rdsRe.Err()
	}

	return rdsRe.Val(), nil
}

func ListSearchAll(key string) (res []string, err error) {
	rdsRe := rdsCli.ZCard(key)
	if nil != rdsRe.Err() {
		return res, rdsRe.Err()
	}

	end := rdsRe.Val()
	if 0 == end {
		return
	}

	rdsRes := rdsCli.ZRange(key, 0, end-1)
	if nil != rdsRes.Err() {
		return res, rdsRes.Err()
	}

	return rdsRes.Val(), nil
}

func KvAdd(key string, val string) (err error) {
	for i := 0; i < 3; i++ {
		err = rdsCli.Set(key, val, 0).Err()
		if nil == err {
			return nil
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("保存kv数据到redis失败：", err, key, val)
	}

	return
}

func KvDel(key string) (err error) {
	for i := 0; i < 3; i++ {
		err = rdsCli.Del(key).Err()
		if nil == err {
			return err
		}
		time.Sleep(RedisFailedRepeatWaitMs * time.Millisecond)
	}

	if nil != err {
		log.Error("删除kv数据到redis失败：", err, key)
	}

	return
}

func KvGet(key string) (string, error) {
	rdsRes := rdsCli.Get(key)
	if redis.Nil == rdsRes.Err() {
		log.Debug("缓存没有数据")
		return "", nil
	} else if nil != rdsRes.Err() {

		return "", rdsRes.Err()
	}

	return rdsRes.Val(), nil
}

func ComputeRange(page, pageSize int) (start, end int64) {
	if MaxPageSize < pageSize {
		pageSize = MaxPageSize
	}

	start = int64((page - 1) * pageSize)
	end = start + int64(pageSize) - 1

	return
}
