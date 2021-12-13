package log

import (
	"fmt"
	"go-websocket-frame/common/global/plugin/config"
	pslog "go-websocket-frame/common/plugin/log"
)

var ftLog *pslog.FluentLog

func Init(c *config.FluentLogConfig) {
	pslog.InitLog(c.Isopen, c.Host,
		c.Port, c.Tag)

	ftLog = pslog.New()
}

func Error(args ...interface{}) {
	ftLog.Write("error", args)
}

func Warn(args ...interface{}) {
	ftLog.Write("warn", args)
}

func Info(args ...interface{}) {
	ftLog.Write("info", args)
}

func Debug(args ...interface{}) {
	ftLog.Write("debug", args)
}

func Errorf(format string, args ...interface{}) {
	ftLog.Write("error", fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	ftLog.Write("warn", fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	ftLog.Write("info", fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...interface{}) {
	ftLog.Write("debug", fmt.Sprintf(format, args...))
}
