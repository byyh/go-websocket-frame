package iface

// 所有独立进程都需要实现该接口
type ProcIface interface {
	Run()        // 运行逻辑函数
	Init()       // 初始化函数
	Destructor() // 析构函数
}
