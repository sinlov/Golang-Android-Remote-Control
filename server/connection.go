package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
)

type connection struct {
	// websocket 连接器
	ws          *websocket.Conn

	messageType chan int
	// 发送信息的缓冲 channel
	message     chan []byte
	// The hub.
	h           *hub
}

func (c *connection) reader() {
	for {
		messageType, message, err := c.ws.ReadMessage()
		// ws library read ping and pong message by SetPingHandler SetPongHandler
		if err != nil {
			break
		}
		c.h.message <- message
		c.h.messageType <- messageType
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.message {
		for messageType := range c.messageType {
			switch messageType {
			case websocket.BinaryMessage:
				fmt.Println("writer binary message", message)
				b_err := c.ws.WriteMessage(websocket.BinaryMessage, message)
				if b_err != nil {
					break
				}
			case websocket.TextMessage:
				fmt.Println("writer text message", message)
				b_err := c.ws.WriteMessage(websocket.TextMessage, message)
				if b_err != nil {
					break
				}
			}
		}

	}
	c.ws.Close()
}

func ping(msg string) error {
	fmt.Println("ping message", msg)
	return nil
}

func pong(msg string) error {
	fmt.Println("pong message", msg)
	return nil
}

var ws_upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
		return true
	}}

type wsHandler struct {
	h *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("RemoteAddr %v, Server Host %v, route %v, Path %v Cookies %v\n",
		req.RemoteAddr, req.Host, req.URL, req.URL.Path, req.Cookies())

	ws, err := ws_upgrader.Upgrade(w, req, nil)
	if err != nil {
		return
	}
	c := &connection{message: make(chan []byte, 256), messageType:make(chan int, 256), ws: ws, h: wsh.h}
	c.h.register <- c
	defer func() {
		c.h.unregister <- c
	}()
	go c.ws.SetPongHandler(pong)
	go c.ws.SetPingHandler(ping)
	go c.writer()
	c.reader()

}
