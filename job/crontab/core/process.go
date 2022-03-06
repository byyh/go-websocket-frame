package core

import (
	"go-websocket-frame/job/crontab/core/iface"
	"go-websocket-frame/job/crontab/internal/svc"
	"sync"

	"github.com/robfig/cron"
	log "github.com/zeromicro/go-zero/core/logx"
)

var (
	routeMaps map[string]*JobStru
	mux       sync.RWMutex
	proc      *Process
)

type JobStru struct {
	Rule string
	Job  iface.ProcIface
}

type JobFunc func(ctx *svc.ServiceContext, data []byte) error

func AddHandleTask(key, rule string, fc iface.ProcIface) {
	mux.Lock()
	defer mux.Unlock()

	if 0 == len(routeMaps) {
		routeMaps = make(map[string]*JobStru)
	}

	routeMaps[key] = &JobStru{
		Rule: rule,
		Job:  fc,
	}
}

type Process struct {
}

func NewProcess() *Process {
	proc = &Process{}

	return proc
}

func (e *Process) Start() {
	log.Info("begin Process")
	c := cron.New() // 新建一个定时任务对象

	for _, v := range routeMaps {

		c.AddJob(v.Rule, v.Job)
	}

	c.Start()

	log.Info("Process start...")
	select {}
}
