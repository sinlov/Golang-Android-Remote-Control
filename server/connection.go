package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
)

type connection struct {
	// websocket 连接器
	ws         *websocket.Conn

	// 发送信息的缓冲 channel
	messageMap chan map[int][]byte
	//message     chan []byte
	// The hub.
	h          *hub
}

func (c *connection) reader() {
	for {
		messageType, message, err := c.ws.ReadMessage()
		// ws library read ping and pong message by SetPingHandler SetPongHandler
		if err != nil {
			break
		}
		mes := make(map[int][]byte)
		mes[messageType] = message
		c.h.messageMap <- mes
		//c.h.message <- message
		//c.h.messageType <- messageType
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.messageMap {
		for key, value := range message {
			switch key {
			case websocket.BinaryMessage:
				fmt.Printf("messageType: %v, message %v\n", key, value)
				b_err := c.ws.WriteMessage(websocket.BinaryMessage, value)
				if b_err != nil {
					break
				}
			case websocket.TextMessage:
				fmt.Printf("messageType: %v, message %v\n", key, value)
				b_err := c.ws.WriteMessage(websocket.TextMessage, value)
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
	c := &connection{messageMap: make(chan map[int][]byte, 256), ws: ws, h: wsh.h}
	c.h.register <- c
	defer func() {
		c.h.unregister <- c
	}()
	go c.ws.SetPongHandler(pong)
	go c.ws.SetPingHandler(ping)
	go c.writer()
	c.reader()

}
