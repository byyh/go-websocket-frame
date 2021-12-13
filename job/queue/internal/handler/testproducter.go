package handler

import (
	"fmt"
	"go-websocket-frame/common/global"
	"go-websocket-frame/common/global/plugin/log"
)

func TestProducter() {
	log.Info("begin producer msg")
	mq := global.GetMq()
	for i := 1; i <= 1000000; i++ {
		mq.Publish("tst_exchange_1", "", []byte(fmt.Sprintf("mess-no-%d-%d", i, 2)))

		// if 0 == i%2 {
		// 	mq.Publish("tst_exchange_1", "", []byte(fmt.Sprintf("mess-no-%d-%d", i, 2)))
		// } else if 0 == i%3 {
		// 	mq.Publish("tst_queue_2", "", []byte(fmt.Sprintf("mess-no-%d-%d", i, 3)))
		// } else {
		// 	mq.Publish("tst_queue_3", "", []byte(fmt.Sprintf("mess-no-%d-%s", i, "other")))
		// }
	}

	log.Info("end producer msg")
}
