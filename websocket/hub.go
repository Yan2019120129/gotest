package main

import "encoding/json"

// hub 定义websocket链接的集线器。
type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

// 管理链接和消息分发
var h = hub{
	c: make(map[*connection]bool),
	b: make(chan []byte),
	r: make(chan *connection),
	u: make(chan *connection),
}

// run 方法是hub结构的方法，用于启用集线器并处理链接注册，注销和消息广播。
func (h *hub) run() {
	for {
		select {
		// 当连接注册时，将连接添加到集线器中的链接映射中，并发送握手消息。
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
			// 有连接注销时，从集线器的连接，映射器中移除链接，并关闭链接的消息，通道。
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
			// 当有消息要广播时，遍历所有链接，并发送消息。
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					// 如果发送失败表明链接已经关闭，从集线器中移除链接并关闭连接的消息通道。
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
