package okx

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
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
				maxReconnect: 5,                // 最大重连次数
				message:      []byte("ping"),
				data:         make(chan string, 10), // 用于接收参数。
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
	data         chan string // 用于接收参数。
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
	fmt.Println("发送信息")
	defer instance.wg.Done() // 在函数退出时递减计数器
	if err := instance.conn.WriteMessage(websocket.TextMessage, instance.message); err != nil {
		panic(err)
		return
	}
}

// ReadMessages 读取消息。
func (instance *okxInstance) ReadMessages() {
	fmt.Println("读取信息")
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
		instance.data <- string(message)
	}
}

// HandleMessages 处理收到的信息。
func (instance *okxInstance) handleMessages() {
	fmt.Println("处理信息")
	defer instance.wg.Done() // 在函数退出时递减计数器
	for {
		message := <-instance.data
		log.Println("收到消息:", message)
	}
}

// heartbeatMessage 测试websocket心跳。
func (instance *okxInstance) heartbeatMessage() {
	fmt.Println("测试心跳")
	defer instance.wg.Done() // 在函数退出时递减计数器
	for {
		// 发送心跳消息
		err := instance.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			// 断开连接的时候重新连接
			for i := 0; i < instance.maxReconnect; i++ {
				fmt.Println("重新连接-", i)
				err := instance.ConnectWS()
				if err != nil {
					continue
				}
				return
			}
		}
		// 每隔三秒发送消息时间
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
