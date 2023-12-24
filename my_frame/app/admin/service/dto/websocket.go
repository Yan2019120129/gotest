package dto

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// Data 用于存储websocket数据
type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

// connection 表示websocket 链接。
type connection struct {
	ws   *websocket.Conn // 底层时wb链接。
	sc   chan []byte     // sc是用于发送消息的通道。
	data *Data           // 包含与链接相关的数据
}

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
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

// 存储当前连接的用户列表。
var user_list = []string{}

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

// writer 负责处理从ws读取的消息。
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

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
