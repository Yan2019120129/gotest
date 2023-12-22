package index

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader       = websocket.Upgrader{ReadBufferSize: 512, WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}
	messagesLock   sync.Mutex
	publicVariable string
)

// SetPublicVariable 设置公共变量的值
func SetPublicVariable(value string) {
	messagesLock.Lock()
	publicVariable = value
	messagesLock.Unlock()
}

// WebsocketServer 处理来自客户端的WebSocket连接。
func WebsocketServer(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// 模拟欢迎消息
	welcomeMessage := "欢迎来到人工客服！请问有什么我可以帮助您的？"
	if err := conn.WriteMessage(websocket.TextMessage, []byte(welcomeMessage)); err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println("进入循环。")
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println("服务端：从客户端收到消息:", string(msg))

		// 在接收到消息后，设置公共变量的值
		SetPublicVariable(string(msg))

		// 在不为空时回复公共变量的信息
		if publicVariable != "" {
			fmt.Println("服务端：放置到公共变量:", publicVariable)
			if err = conn.WriteMessage(messageType, []byte(publicVariable)); err != nil {
				log.Println(err)
				return nil, err
			}
			publicVariable = ""
		}
	}
}

// WebsocketClient 处理WebSocket连接作为客户端。
func WebsocketClient(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	fmt.Println("进入循环。")
	for {
		fmt.Println("开始执行。")
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println("客户端：从服务器收到消息:", string(msg))

		// 在不为空时回复公共变量的信息
		if publicVariable != "" {
			fmt.Println("客户端：设置公共变量:", publicVariable)
			if err = conn.WriteMessage(messageType, []byte(publicVariable)); err != nil {
				log.Println(err)
				return nil, err
			}
			publicVariable = ""
		}
	}
}
