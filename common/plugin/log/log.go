/*
 * 日志采集对接
 * 接入单个或多个数据库均需要在这里配置
 *
 * 注意：目前数据库初始化包括 gorm 库的初始化，以及 beego 的orm 初始化
 * 在调用的时候可以根据需要调用相关的库
 * 调用gorm的时候请采用 GormDb() 获取db，db会自动
 */

package log

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/byyh/go/com"
	"github.com/fluent/fluent-logger-golang/fluent"
)

type FluentLog struct {
	Logger *fluent.Fluent
	Tag    string
}

var (
	ftLog       *FluentLog
	isFluentLog bool
)

func New() *FluentLog {
	return ftLog
}

func InitLog(isFluentLog bool, host string, port int, tag string) {
	var err error

	if !isFluentLog {
		return
	}

	log.Println("begin init log")
	ftLog = &FluentLog{
		Tag: tag,
	}

	ftLog.Logger, err = fluent.New(fluent.Config{
		FluentPort: port,
		FluentHost: host,
		MaxRetry:   100,
		Async:      false,
	})
	if nil != err {
		log.Println("init fluent log failed", err)
	}
}

func (this *FluentLog) Error(args ...interface{}) {
	this.Write("error", args)
}

func (this *FluentLog) Warn(args ...interface{}) {
	this.Write("warn", args)
}

func (this *FluentLog) Info(args ...interface{}) {
	this.Write("info", args)
}

func (this *FluentLog) Debug(args ...interface{}) {
	this.Write("debug", args)
}

func (this *FluentLog) Write(level string, in ...interface{}) {
	var (
		args []interface{}
		hStr string
	)

	if 0 >= len(in) {
		log.Println("没有日志需要输出", level)
		return
	}

	_, file, line, _ := runtime.Caller(2)
	hStr = fmt.Sprintf("file:%s,line:%d,", file, line)
	args = append(args, hStr, ", ", level)
	for _, v := range in {
		args = append(args, v)
	}

	if !isFluentLog {
		log.Println(args...)

		return
	}

	//log.Println("log-p=", level, args)
	if nil == ftLog || nil == ftLog.Logger {
		panic(errors.New("日志未初始化"))
	}
	tag := this.Tag + "-" + level

	mp := make(map[string]interface{})
	mp["tag"] = tag
	mp["level"] = level
	mp["log_time"] = new(com.Time).Now()

	var newArgs []string
	num := len(args)
	for i := 0; i < num; i++ {
		tmp := args[i].([]interface{})
		for _, arg := range tmp {
			var str string
			if "string" != com.Typeof(arg) {
				str = fmt.Sprintf("%s", arg)
			} else {
				str = arg.(string)
			}

			newArgs = append(newArgs, str)
		}
	}

	mp["data"] = newArgs

	go this.Post(tag, mp)
}

func (this *FluentLog) Post(tag string, mp map[string]interface{}) {
	defer func() {
		if err := recover(); nil != err {
			log.Println("fluent lib occur exception: ", err)
		}
	}()
	//log.Println("post data===", tag, mp)
	err := this.Logger.Post(tag, mp)
	if nil != err {
		log.Println("post fluent log failed", err)
		log.Println("err-data===", mp["data"])
	}
	//log.Println("发送完毕", mp["data"])
}
