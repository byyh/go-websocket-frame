// 独立进程适配
// 新增独立进程处理结构体，不用修改该文件，只需要处理结构体实现iface.ProcIface接口

package core

import (
	"runtime/debug"
	"sync"

	"go-websocket-frame/job/crontab/core/iface"

	"github.com/byyh/go/com"
	log "github.com/zeromicro/go-zero/core/logx"
)

var (
	runLock sync.Mutex

	runProcList map[string]bool // 保障进程同一时间只有一个实例在跑
)

func init() {
	runProcList = make(map[string]bool)
}

type runProcess struct {
	Job iface.ProcIface
}

func NewRunProcess(job iface.ProcIface) *runProcess {
	return &runProcess{
		Job: job,
	}
}

func (p *runProcess) Run() {
	defer p.catchError("启动进程")

	runLock.Lock()
	// 避免同时运行多个
	if _, ok := runProcList[com.Typeof(p.Job)]; ok {
		runLock.Unlock()
		log.Info(com.Typeof(p.Job), " is running...", new(com.Time).Now())
		return
	}
	runProcList[com.Typeof(p.Job)] = true
	runLock.Unlock()

	log.Info(com.Typeof(p.Job), " is start...", new(com.Time).Now())

	p.call(p.Job)

	delete(runProcList, com.Typeof(p.Job))
	log.Info(com.Typeof(p.Job), " run over...", new(com.Time).Now())
}

func (p *runProcess) call(job iface.ProcIface) {
	defer p.catchError("调用Run方法")

	job.Init()
	job.Run()
	job.Destructor()
}

func (p *runProcess) catchError(action string) {
	if err := recover(); nil != err {
		delete(runProcList, com.Typeof(p.Job))

		// 写错误异常日志
		log.Error("exec ", action, " error: handle recover", err)
		log.Error("异常栈：", string(debug.Stack()))
	}
}
