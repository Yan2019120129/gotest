package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// connection 表示websocket 链接。
type connection struct {
	ws   *websocket.Conn // 底层时wb链接。
	sc   chan []byte     // sc是用于发送消息的通道。
	data *Data           // 包含与链接相关的数据
}

// wu 是websocket升级器，负责将HTTP链接升级微websocket 链接。
var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

// myws 连接http的处理器。
func myws(w http.ResponseWriter, r *http.Request) {
	// 将HTTP链接升级为websocket 链接。
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// 创建结构体，并将其注册到集线器中。
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	h.r <- c

	// 启动连接器，写入协程。
	go c.writer()

	// 处理链接的读取。
	c.reader()

	// 连接关闭时执行的清理工作。
	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

// writer 负责处理从ws读取的消息。
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

// 存储当前连接的用户列表。
var user_list = []string{}

// reader 读取连接消息的方法。
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			// 处理登录消息，更新用户列表，并广登录消息。
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			// 处理用户消息，广播用户消息。
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			// 处理用户推出登录消息，更新用户列表，并广播消息。
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			// 处理未知消息。
			fmt.Print("========default================")
		}
	}
}

// del 从切片中删除指定元素。
func del(slice []string, user string) []string {
	count := len(slice)
	// 如果没有用户信息直接返回。
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
