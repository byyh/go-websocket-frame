package ws

import (
	protomod "go-websocket-frame/api/ws/proto"
	"go-websocket-frame/common/atom"
	"time"

	"go-websocket-frame/common/utils"
	"sync"

	"go-websocket-frame/common/global/plugin/log"

	"github.com/golang/protobuf/proto"
)

var (
	routeMaps map[uint32]DispatchFunc
	mux       sync.RWMutex
)

type DispatchFunc func(c *Client, cmd uint32, reqUid int64, data []byte) atom.Error

func AddWsHandleFunc(cmd uint32, fc DispatchFunc) {
	mux.Lock()
	defer mux.Unlock()

	if 0 == len(routeMaps) {
		routeMaps = make(map[uint32]DispatchFunc)
	}
	log.Info("addwsroute: ", cmd)
	routeMaps[cmd] = fc
}

func dispatch(c *Client, message []byte) {
	log.Info("dispatch-begin:", string(message))
	// 捕获异常
	defer utils.CatchExceptionFunc("ws-dispatch", func() {
		log.Error("ws.dispatch-处理<", message, ">发生异常,excption-cli-uid=", c.User)
	})

	msg, err := decodeProto(message)
	if nil != err {
		log.Error("ws.dispatch-json解析<", message, ">发生错误,", err)

		return
	}

	if h, ok := routeMaps[msg.Cmd]; ok {
		if err := h(c, msg.Cmd, msg.ReqUid, []byte(msg.Data)); nil != err {
			log.Error("执行<", msg.Cmd, ">失败,cli-uid=", c.User.Id, ",", err)
		}
	} else {
		log.Error("ws客户端上传cmd非法,", msg.Cmd, ",cli-uid=", c.User.Id)
	}
}

func encodeProto(cmd uint32, reqUid int64, data proto.Message) ([]byte, error) {
	bt, err := proto.Marshal(data)
	if nil != err {
		log.Errorf("Client.SendByte-encodeProto-err,%s,%s", data, err)
		return bt, err
	}

	msg := &protomod.Msg{
		Cmd:    (cmd),
		Data:   bt,
		ReqUid: reqUid,
		Tms:    time.Now().UnixNano() / 1e6,
	}

	btMsg, err := proto.Marshal(msg)
	if nil != err {
		log.Error("解压上传proto数据失败")
	}

	return btMsg, err
}

func decodeProto(bt []byte) (*protomod.Msg, error) {
	msg := &protomod.Msg{}

	if err := proto.Unmarshal(bt, msg); nil != err {
		log.Error("解压上传proto数据失败")
		return nil, err
	}

	return msg, nil
}
