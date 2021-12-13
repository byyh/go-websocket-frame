package cache

import (
	"sync"

	"github.com/go-redis/redis/v7"
)

const Nil = redis.Nil

var (
	Redis        *redis.Client
	ClusterRedis *redis.ClusterClient
	initRdsLock  sync.Mutex
)

func New() *redis.Client {
	return Redis
}

func NewCluster() *redis.ClusterClient {
	return ClusterRedis
}

func InitRedisCluster(addrs []string) {
	initRdsLock.Lock()
	defer initRdsLock.Unlock()

	ClusterRedis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		// 可选的密码。
		// Password: redisConfig.GetPassword(),
	})
}

func InitRedisSingle(addr, passwd string, dbNum int) {
	initRdsLock.Lock()
	defer initRdsLock.Unlock()

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       dbNum,
	})
}
