// 客户端

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
)

// Client 用于发送信息
func Client() {
	serverAddr := "ws://localhost:8080/websocket"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u, _ := url.Parse(serverAddr)
	q := u.Query()
	q.Set("token", "your_token")
	u.RawQuery = q.Encode()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("WebSocket 连接错误:", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})

	// 启动 goroutine 用于接收服务端消息
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("客户端接收消息错误:", err)
				return
			}
			fmt.Printf("客户端收到消息: %s\n", message)
		}
	}()

	// 向服务端发送消息
	message := []byte("Hello, Server!")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("客户端发送消息错误:", err)
		return
	}

	select {
	case <-done:
	case <-interrupt:
		fmt.Println("接收到中断信号，关闭连接...")
		return
	}
}
