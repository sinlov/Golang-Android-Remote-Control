package server

type hub struct {
	// 注册了的连接器
	connections map[*connection]bool

	// 消息类型
	messageType chan int
	// 从连接器中发入的各种信息
	message     chan []byte

	// 从连接器中注册请求
	register    chan *connection

	// 从连接器中注销请求
	unregister  chan *connection
}

func newHub() *hub {
	return &hub{
		connections:        make(map[*connection]bool),
		message:make(chan []byte),
		messageType: make(chan int),
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
				close(c.message)
			}
		case t := <-h.messageType:
			for c := range h.connections {
				select {
				case c.messageType <- t:
				}
			}
		case m := <-h.message:
			for c := range h.connections {
				select {
				case c.message <- m:
				default:
					delete(h.connections, c)
					close(c.message)
				}
			}

		}
	}
}

