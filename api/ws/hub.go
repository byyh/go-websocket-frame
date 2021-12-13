// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"go-websocket-frame/common/global/plugin/log"
	"go-websocket-frame/common/utils"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients sync.Map

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// 注销所有客户端
	unregisterAll chan bool
}

var (
	hub *Hub
)

func NewHub() *Hub {
	hub = &Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		unregisterAll: make(chan bool),
	}

	return hub
}

func (h *Hub) Run() {
	for {
		defer utils.CatchException("hub-for-select")

		select {
		// 注册
		case client := <-h.register:
			log.Info("注册", client)
			h.clients.Store(client, true)
		// 注销
		case client := <-h.unregister:
			if _, ok := h.clients.Load(client); ok {
				log.Info("注销：", client)
				close(client.send)
				close(client.recv)
				h.clients.Delete(client)
			}
			// 批量发通知
		case message := <-h.broadcast:
			h.clients.Range(func(k, v interface{}) bool {
				client := k.(*Client)

				select {
				case client.send <- message:
				default:
					close(client.send)
					h.clients.Delete(client)
				}

				return true
			})
		// 全部注销
		case <-h.unregisterAll:
			h.clients.Range(func(k, v interface{}) bool {
				client := k.(*Client)

				close(client.send)
				h.clients.Delete(client)

				return true
			})

		}
	}
}

// 外部全部注销调用
func CloseAllClient() {
	hub.unregisterAll <- true
}
