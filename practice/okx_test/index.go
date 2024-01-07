package okx_test

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gotest/practice/okx_test/dto"
	"log"
	"sync"
)

const ServerAddr = "wss://ws.okx.com:8443/ws/v5/public"

// Instance websocket实例
type Instance struct {
	wg   sync.WaitGroup
	conn *websocket.Conn
	Data chan string
}

func main() {
	ws := ConnectWS(ServerAddr)
	defer ws.conn.Close()

	params := &dto.TickerParams{
		Op: "subscribe",
		Args: []dto.Arg{
			{Channel: "tickers", InstID: "MDT-USDT"},
		},
	}

	ws.SendMessages(params)
	data := &dto.TickerData{}

	go ws.ReadMessages(data)
	//ws.wg.Add(1)
	go ws.HandleMessages()
	ws.wg.Add(2)
	ws.wg.Wait()
}

// ConnectWS 连接 websocket
func ConnectWS(serverAddr string) *Instance {
	// 创建一个 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("WebSocket连接错误:", err)
	}

	return &Instance{
		conn: conn,
		Data: make(chan string, 30),
	}
}

// SendMessages 发送消息
func (instance *Instance) SendMessages(params interface{}) {
	// 将 Message 结构体序列化为 JSON 字符串
	jsonData, err := json.Marshal(params)
	if err != nil {
		log.Fatal("JSON序列化错误:", err)
		return
	}

	// 使用 WebSocket 连接发送 JSON 数据
	err = instance.conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Fatal("消息发送错误:", err)
		return
	}

	fmt.Println("JSON 数据已发送成功！")
}

// ReadMessages 读取消息
func (instance *Instance) ReadMessages(data interface{}) {
	for {
		_, message, err := instance.conn.ReadMessage()
		if err != nil {
			log.Println("消息接收错误:", err)
			break
		}

		if err := json.Unmarshal(message, data); err != nil {
			log.Println("JSON解析错误:", err)
			continue
		}

		instance.Data <- string(message)
		fmt.Println("读取的数据:", string(message))
	}
}

// HandleMessages 处理收到的信息
func (instance *Instance) HandleMessages() {
	// 主 goroutine 用于处理接收到的消息
	for {
		message := <-instance.Data
		fmt.Println("收到消息:", message)
	}
}
