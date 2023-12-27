package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type TickerParams struct {
	Op   string `json:"op"`
	Args []Arg  `json:"args"`
}

type TickerData struct {
	Arg  Arg     `json:"arg"`
	Data []Datum `json:"data"`
}

type Arg struct {
	Channel string `json:"channel"`
	InstID  string `json:"instId"`
}

type Datum struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
}

const ServerAddr = "wss://ws.okx.com:8443/ws/v5/public"

var wg sync.WaitGroup

// Instance websocket实例
type Instance struct {
	conn *websocket.Conn
	Data chan string
}

func main() {
	ws := ConnectWS(ServerAddr)
	defer ws.conn.Close()

	params := &TickerParams{
		Op: "subscribe",
		Args: []Arg{
			{Channel: "tickers", InstID: "MDT-USDT"},
		},
	}

	ws.SendMessages(params)
	data := &TickerData{}

	go ws.ReadMessages(data)
	wg.Add(1)
	go ws.HandleMessages()
	wg.Add(1)
	wg.Wait()
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
