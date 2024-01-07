package ws

import (
	"encoding/json"
	"fmt"
)

// Hub 定义websocket链接的集线器。
type Hub struct {
	c map[*Connection]bool
	B chan []byte
	R chan *Connection
	u chan *Connection
}

// H 管理链接和消息分发
var H = Hub{
	c: make(map[*Connection]bool),
	u: make(chan *Connection),
	B: make(chan []byte),
	R: make(chan *Connection),
}

// User_list 存储当前连接的用户列表。
var User_list []string

// Run 方法是hub结构的方法，用于启用集线器并处理链接注册，注销和消息广播。
func (h *Hub) Run() {
	for {
		select {
		// 当连接注册时，将连接添加到集线器中的链接映射中，并发送握手消息。
		case c := <-h.R:
			h.c[c] = true
			c.Data.Ip = c.Ws.RemoteAddr().String()
			c.Data.Type = "handshake"
			c.Data.UserList = User_list
			data_b, _ := json.Marshal(c.Data)
			c.Sc <- data_b
			// 有连接注销时，从集线器的连接，映射器中移除链接，并关闭链接的消息，通道。
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.Sc)
			}
			// 当有消息要广播时，遍历所有链接，并发送消息。
		case data := <-h.B:
			for c := range h.c {
				select {
				case c.Sc <- data:
				default:
					// 如果发送失败表明链接已经关闭，从集线器中移除链接并关闭连接的消息通道。
					delete(h.c, c)
					close(c.Sc)
				}
			}
		}
	}
}

// Del 从切片中删除指定元素。
func Del(slice []string, user string) []string {
	// 如果没有用户信息直接返回。
	count := len(slice)
	if count == 0 {
		return slice
	}

	// 如果只有一条信息或者这条信息刚好是要查找的信息，则删除，返回空数组。
	if count == 1 && slice[0] == user {
		return []string{}
	}

	// 新建数组，将删除用户后的数据放入数组。
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
