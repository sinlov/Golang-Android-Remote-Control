package server

import "fmt"

type hub struct {
	// 注册了的连接器
	connections map[*connection]bool

	// 从连接器中发入的各种信息
	messageMap  chan map[int][]byte
	//message     chan []byte

	// 从连接器中注册请求
	register    chan *connection

	// 从连接器中注销请求
	unregister  chan *connection
}

func newHub() *hub {
	return &hub{
		connections:        make(map[*connection]bool),
		messageMap: make(chan map[int][]byte),
		register:        make(chan *connection),
		unregister:        make(chan *connection),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.messageMap)
			}
		case t := <-h.messageMap:
			for c := range h.connections {
				select {
				case c.messageMap <- t:
				default:
					fmt.Printf("delete messageType %v", c.messageMap)
					delete(h.connections, c)
					close(c.messageMap)
				}
			}
		}
	}
}

