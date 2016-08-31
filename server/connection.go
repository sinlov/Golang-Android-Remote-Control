package server

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	// websocket 连接器
	ws   *websocket.Conn

	// 发送信息的缓冲 channel
	send chan []byte

	// The hub.
	h    *hub
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
		b_err := c.ws.WriteMessage(websocket.BinaryMessage, message)
		if b_err != nil {
			break
		}
	}
	c.ws.Close()
}

var ws_upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
		return true
	}}

type wsHandler struct {
	h *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := ws_upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws, h: wsh.h}
	c.h.register <- c
	defer func() {
		c.h.unregister <- c
	}()
	go c.writer()
	c.reader()
}
