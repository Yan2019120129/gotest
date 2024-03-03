package indexserver

import (
	"encoding/json"
	ws "gotest/frame/my_frame/module/websocket"
	"net/http"
)

// WebsocketServer 处理来自客户端的WebSocket连接。
func WebsocketServer(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	// 将HTTP链接升级为websocket 链接。
	socket, err := ws.WU.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
		return
	}

	// 创建结构体，并将其注册到集线器中。
	c := &ws.Connection{Sc: make(chan []byte, 256), Ws: socket, Data: &ws.Data{}}
	ws.H.R <- c

	// 启动连接器，写入协程。
	go c.Writer()

	// 处理链接的读取。
	c.Reader()
	// 连接关闭时执行的清理工作。
	defer func() {
		c.Data.Type = "logout"
		ws.User_list = ws.Del(ws.User_list, c.Data.User)
		c.Data.UserList = ws.User_list
		c.Data.Content = c.Data.User
		data_b, _ := json.Marshal(c.Data)
		ws.H.B <- data_b
		ws.H.R <- c
	}()

	return "ok", nil
}
