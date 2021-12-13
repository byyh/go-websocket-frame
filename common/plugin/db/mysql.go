/*
 * 文件用于接入数据库，所有数据库初始化在这个文件完成
 * 接入单个或多个数据库均需要在这里配置
 *
 * 注意：目前数据库初始化包括 gorm 库的初始化
 *
 *
 * 调用gorm的时候请采用 db.NewDb() 获取db，db会自动处理长链接断开的问题
 */

package db

import (
	"errors"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormDb     *gorm.DB
	initDbLock sync.Mutex
)

func New() *gorm.DB {

	return gormDb
}

// 初始化
func InitDb(dns string) {
	initDbLock.Lock()
	defer initDbLock.Unlock()

	if "" == dns {
		panic(errors.New("config.MysqlDns.Default is empty"))
		return
	}
	gormDb = GormDb(dns)
	SetDbConfig(gormDb)
}

func SetDbConfig(ob *gorm.DB) {
	gdb, _ := ob.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	gdb.SetMaxIdleConns(100)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	gdb.SetMaxOpenConns(1000)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	gdb.SetConnMaxLifetime(time.Hour)

}

// gorm 调用必须通过这个方法调用,函数处理了连接池和断开自动重连接的处理。
func GormDb(sqlDns string) *gorm.DB {
	gormDb, err := gorm.Open(mysql.Open(sqlDns),
		&gorm.Config{})
	if nil != err {
		panic("gorm数据库连接错误！" + sqlDns + err.Error())
	}

	//gdb, _ := gormDb.DB()

	return gormDb
}
