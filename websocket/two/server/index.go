// 服务端

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket 升级错误:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("消息读取错误:", err)
			return
		}

		fmt.Printf("服务端收到消息: %s\n", p)

		// 向客户端发送消息
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("服务端发送消息错误:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/websocket", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
