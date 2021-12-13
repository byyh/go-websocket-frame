package handler

import (
	"go-websocket-frame/common/global/plugin/queue"
)

func InitRoute(cr *queue.Consumer) {
	// 参数说明： 队列名称   是否自动应答  消费逻辑函数
	cr.AddConsumerHandler("tst_queue_1", false, TestHandler1) // demo示例
	cr.AddConsumerHandler("tst_queue_2", false, TestHandler2) // demo示例
	cr.AddConsumerHandler("tst_queue_3", false, TestHandler3) // demo示例
}
