// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"bytes"
	"context"
	"errors"
	"go-websocket-frame/api/internal/svc"
	"go-websocket-frame/api/internal/types"
	proto "go-websocket-frame/api/ws/proto"
	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-websocket-frame/common/global/plugin/log"

	sproto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 20 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 5120
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	// CheckOrigin: func(r *http.Request) bool {
	// 	return true
	// },
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
	recv chan []byte

	// 用户信息,一个客户端对应一个用户
	User *types.UserBase

	SvcCtx *svc.ServiceContext

	Ctx context.Context
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Info("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// 调用ws路由分发
		dispatch(c, message)
		//c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(ctx *gin.Context, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer utils.CatchException("ServeWs异常了：")

	v, ok := ctx.Get("user_info")
	if !ok {
		panic(errors.New("用户未设置"))
	}
	log.Info("----------------", v.(*types.UserBase).Name)
	client := &Client{Ctx: r.Context(), hub: hub, conn: conn,
		send: make(chan []byte, maxMessageSize), User: v.(*types.UserBase)}

	log.Info("---------client-------", client)
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

// 发送消息给客户端
func (c *Client) SendMsg(cmd uint32, reqUid int64, data sproto.Message) {
	defer utils.CatchExceptionFunc("Client.SendByte", func() {
		log.Errorf("Client.SendByte-异常,发送的数据：%d,%s ", cmd, data)
	})

	btMsg, err := encodeProto(cmd, reqUid, data)
	if nil != err {
		log.Errorf("Client.SendByte-encodeProto-err,%s,%s", data, err)
		return
	}

	c.send <- btMsg
}

// 发送消息给所有客户端（广播）
func (c *Client) SendBroadcast(cmd uint32, data sproto.Message) {
	defer utils.CatchExceptionFunc("Client.SendBroadcast", func() {
		log.Errorf("Client.SendBroadcast-异常,发送的数据：%d,%s ", data)
	})

	btMsg, err := encodeProto(cmd, 0, data)
	if nil != err {
		log.Errorf("Client.SendBroadcast-encodeProto-err,%s,%s", data, err)
		return
	}

	hub.broadcast <- btMsg
}

// 发生错误给客户端
func (c *Client) SendErrorMsg(cmd uint32, reqUid int64, errcode int) {
	defer utils.CatchExceptionFunc("Client.SendErrorMsg", func() {
		log.Errorf("Client.SendErrorMsg-异常,发送的数据：%d,%d,%d ", cmd, reqUid, errcode)
	})

	data := &proto.ErrResp{
		Code:      int32(errcode),
		Msg:       atom.GetMsgByCode(errcode),
		OriginCmd: cmd,
	}

	btMsg, err := encodeProto(uint32(proto.CmdErr_RetErr), reqUid, data)
	if nil != err {
		log.Errorf("Client.SendErrorMsg-encodeProto-err,%s,%s", data, err)
		return
	}

	c.send <- btMsg
}

func (c *Client) GetUser() *types.UserBase {
	return c.User
}
