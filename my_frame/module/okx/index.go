package okx

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"gotest/my_frame/module/cache"
	"gotest/my_frame/utils"
	"log"
	"sync"
	"time"
)

// ServerTickerAddr 产品行情地址。
const ServerTickerAddr = "wss://ws.okx.com:8443/ws/v5/public"

// ServerCandleAndTradeAddr 产品k线图&全部交易频道地址
const ServerCandleAndTradeAddr = "wss://ws.okx.com:8443/ws/v5/business"

// ServerAddr 存连接频道的服务地址。
var ServerAddr = map[string]string{
	"tickers":       ServerTickerAddr,
	"option-trades": ServerTickerAddr,
	"trades":        ServerCandleAndTradeAddr,
	"trades-all":    ServerCandleAndTradeAddr,
}

type Message struct {
	Arg  Arg    `json:"arg"`
	Data []Data `json:"data"`
}

type Arg struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type Data struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	Ts        string `json:"ts"`
}

var OkxInstance *okxInstance

// 用于保证websocket单例
var _once sync.Once

// init 初始化oke
func init() {
	if OkxInstance == nil {
		_once.Do(func() {
			OkxInstance = &okxInstance{
				wg:           new(sync.WaitGroup),
				conn:         new(websocket.Conn),
				serverAddr:   ServerTickerAddr, // 连接地址
				maxReconnect: 2,                // 最大重连次数
				//message:      []byte("ping"),
				message: []byte("{\n    \"op\": \"subscribe\",\n    \"args\": [\n        {\n            \"channel\": \"tickers\",\n            \"instId\": \"MDT-USDT\"\n        },\n        {\n            \"channel\": \"tickers\",\n            \"instId\": \"1INCH-EUR\"\n        }\n    ]\n}"),
				data:    make(chan []byte, 10), // 用于接收参数。
			}
		})
	}
}

// okxInstance 连接okxwebsocket 实例
type okxInstance struct {
	wg           *sync.WaitGroup
	conn         *websocket.Conn
	serverAddr   string      // 连接地址
	maxReconnect int         // 最大重连次数
	publishName  string      // 生产的消息名
	message      []byte      // 用于发送的数据
	data         chan []byte // 用于接收参数。
}

// ConnectWS 连接okx websocket。
func (instance *okxInstance) ConnectWS() (err error) {
	if instance.serverAddr == "" {
		panic(errors.New("服务地址不能为空！！！"))
	}
	instance.conn, _, err = websocket.DefaultDialer.Dial(instance.serverAddr, nil)
	if err != nil {
		go instance.heartbeatMessage()
		instance.wg.Add(1)
	}
	go instance.SendMessages()
	instance.wg.Add(1)
	go instance.ReadMessages()
	instance.wg.Add(1)
	go instance.handleMessages()
	instance.wg.Add(1)
	instance.wg.Wait()
	return err
}

// SendMessages 发送消息。
func (instance *okxInstance) SendMessages() {
	defer fmt.Println("发送信息关闭")
	defer instance.wg.Done() // 在函数退出时递减计数器
	if err := instance.conn.WriteMessage(websocket.TextMessage, instance.message); err != nil {
		panic(err)
		return
	}
}

// ReadMessages 读取消息。
func (instance *okxInstance) ReadMessages() {
	defer fmt.Println("读取信息关闭")
	defer instance.wg.Done() // 在函数退出时递减计数器
	for {
		_, message, err := instance.conn.ReadMessage()
		if err != nil {
			log.Println("消息接收错误:", err)
			go instance.heartbeatMessage()
			instance.wg.Add(1)
			break
		}
		log.Println("读取到的消息:", string(message))
		instance.data <- message
	}
}

// HandleMessages 处理收到的信息。
func (instance *okxInstance) handleMessages() {
	defer fmt.Println("处理信息关闭")
	defer instance.wg.Done() // 在函数退出时递减计数器
	for {
		message := <-instance.data
		log.Println("收到消息:", string(message))
		data := &Message{}
		utils.ByteListToObj(message, data)
		cache.Publish(data.Arg.Channel+"-"+data.Arg.InstId, utils.ObjToByteList(data.Data))
	}
}

// heartbeatMessage 测试websocket心跳。
func (instance *okxInstance) heartbeatMessage() {
	defer fmt.Println("测试心跳关闭")
	defer instance.wg.Done() // 在函数退出时递减计数器
	for {
		// 发送心跳消息
		err := instance.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			// 断开连接的时候重新连接
			for i := 0; i < instance.maxReconnect; i++ {
				instance.maxReconnect--
				fmt.Println("重新连接-", instance.maxReconnect)
				err := instance.ConnectWS()
				if err != nil {
					continue
				}

				if err := instance.conn.Close(); err != nil {
					defer instance.wg.Done() // 在函数退出时递减计数器
					return
				}
				return
			}
		}
		// 每隔三秒发送消息
		time.Sleep(3 * time.Second)
	}
}

// SetServerAddr 设置服务器地址
func (instance *okxInstance) SetServerAddr(serverAddr string) {
	instance.serverAddr = serverAddr
}

// SetMaxReconnect 设置最大重连次数
func (instance *okxInstance) SetMaxReconnect(maxReconnect int) {
	instance.maxReconnect = maxReconnect
}

// SetPublishName 设置发布名字
func (instance *okxInstance) SetPublishName(publishName string) {
	instance.publishName = publishName
}

// SetSendMessage 设置要发送的数据
func (instance *okxInstance) SetSendMessage(message []byte) {
	instance.message = message
}