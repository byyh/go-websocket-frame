// 定义的三方初始化变量，各个业务层可直接使用

package global

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"

	squeue "go-websocket-frame/common/plugin/queue"
)

func IsRedisCluster() bool {
	return redisIsCluster
}

func GetRedis() *redis.Client {
	return redisCli
}

func GetRedisCluster() *redis.ClusterClient {
	return redisClusterCli
}

func GetDb() *gorm.DB {
	return mysqlCli
}

func GetMq() squeue.RabbitmqIface {
	return mqProducer.GetMq()
}
