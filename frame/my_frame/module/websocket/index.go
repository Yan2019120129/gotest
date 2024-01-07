package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// WU 是websocket升级器，负责将HTTP链接升级微websocket 链接。
var WU = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

// Connection 表示websocket 链接。
type Connection struct {
	Ws   *websocket.Conn // 底层时wb链接。
	Sc   chan []byte     // sc是用于发送消息的通道。
	Data *Data           // 包含与链接相关的数据
}

// Writer 负责处理从ws读取的消息。
func (c *Connection) Writer() {
	for message := range c.Sc {
		if err := c.Ws.WriteMessage(websocket.TextMessage, message); err != nil {
			panic(err)
			return
		}
	}
	if err := c.Ws.Close(); err != nil {
		return
	}
}

// Reader 读取连接消息的方法。
func (c *Connection) Reader() {
	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			H.R <- c
			break
		}
		if err = json.Unmarshal(message, &c.Data); err != nil {
			return
		}
		fmt.Println("message：", c.Data)
		switch c.Data.Type {
		case "login":
			// 处理登录消息，更新用户列表，并广登录消息。
			c.Data.User = c.Data.Content
			c.Data.From = c.Data.User
			User_list = append(User_list, c.Data.User)
			c.Data.UserList = User_list
			data_b, _ := json.Marshal(c.Data)
			H.B <- data_b
		case "user":
			// 处理用户消息，广播用户消息。
			c.Data.Type = "user"
			data_b, _ := json.Marshal(c.Data)
			H.B <- data_b
		case "logout":
			// 处理用户推出登录消息，更新用户列表，并广播消息。
			c.Data.Type = "logout"
			User_list = Del(User_list, c.Data.User)
			data_b, _ := json.Marshal(c.Data)
			H.B <- data_b
			H.R <- c
		default:
			// 处理未知消息。
			fmt.Print("========default================")
		}
	}
}
