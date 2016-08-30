package server

type hub struct {
	// 注册了的连接器
	connections map[*connection]bool

	// 从连接器中发入的信息
	broadcast   chan []byte

	// 从连接器中注册请求
	register    chan *connection

	// 从连接器中注销请求
	unregister  chan *connection
}

func newHub() *hub {
	return &hub{
		connections:        make(map[*connection]bool),
		broadcast:                make(chan []byte),
		register:                make(chan *connection),
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
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}

