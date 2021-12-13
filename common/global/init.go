//所有的插件初始化文件

package global

import (
	"go-websocket-frame/common/global/plugin/config"
	"go-websocket-frame/common/global/plugin/log"
	"go-websocket-frame/common/global/plugin/queue"
	"go-websocket-frame/common/plugin/cache"
	"go-websocket-frame/common/plugin/db"
	"strings"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

var (
	// redisCli
	redisIsCluster  bool
	redisCli        *redis.Client
	redisClusterCli *redis.ClusterClient

	// mysql
	mysqlCli *gorm.DB

	// rabbitmq
	mqProducer *queue.Producer
	mqConsumer *queue.Consumer
)

func InitRedis(c *config.RedisConf) {
	if "cluster" == c.Type {
		arr := strings.Split(c.Host, ",")
		cache.InitRedisCluster(arr)

		redisClusterCli = cache.NewCluster()
	} else {
		cache.InitRedisSingle(c.Host, c.Pass, c.Dbnum)

		redisCli = cache.New()
	}

}

// 预留适配
type DbIface interface {
}

func InitMysql(dns string) {
	db.InitDb(dns)
	mysqlCli = db.New()

}

func InitProducer(c *config.RabbitmqConfig) {
	mqProducer = queue.NewProducer(c)
	mqProducer.Init()
}

func InitConsumer(c []*config.RabbitmqConfig) *queue.Consumer {
	mqConsumer = queue.NewConsumer(c)

	return mqConsumer
}

func InitLog(c *config.FluentLogConfig) {
	log.Init(c)
}
