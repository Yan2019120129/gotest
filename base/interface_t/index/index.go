package index

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"strconv"
	"time"
)

var message = []byte(`{
	"op": "subscribe",
	"data": [
		[
            "1708617300000",
            "51573.7",
            "51621.8",
            "51573.7",
            "51602.5",
            "32.19050846",
            "1661243.0761018",
            "1661243.0761018",
            "1"
        ],
        [
            "1708617000000",
            "51531.2",
            "51576.9",
            "51506.1",
            "51573.6",
            "37.36719661",
            "1926220.73216047",
            "1926220.73216047",
            "1"
        ]
	]
}`)

// Message 消息体
type Message struct {
	Op   string        `json:"op"`   //	方法名称
	Data []interface{} `json:"data"` //	方法参数
}

// KlineData k线图数据
type KlineData struct {
	OpenPrice  float64 `json:"openPrice"`  //开盘价格
	HighPrice  float64 `json:"highPrice"`  //最高价格
	LowsPrice  float64 `json:"lowsPrice"`  //最低价格
	ClosePrice float64 `json:"closePrice"` //收盘价格
	Vol        float64 `json:"vol"`        //交易量
	Amount     float64 `json:"amount"`     //成交额
	CreatedAt  int64   `json:"createdAt"`  //开盘时间
}

// InterfaceToObj interface转换结构体类型
func InterfaceToObj(data interface{}, obj interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataBytes, obj)
	return err
}

// ToKline 二维字符串数组转换为KlineData
func ToKline(kline []string) (data *KlineData) {
	createTime, err := strconv.ParseInt(kline[0], 10, 64)
	if err != nil {
		panic(err)
	}

	openPrice, err := strconv.ParseFloat(kline[1], 64)
	if err != nil {
		panic(err)
	}

	highPrice, err := strconv.ParseFloat(kline[2], 64)
	if err != nil {
		panic(err)
	}

	lowsPrice, err := strconv.ParseFloat(kline[3], 64)
	if err != nil {
		panic(err)
	}

	closePrice, err := strconv.ParseFloat(kline[4], 64)
	if err != nil {
		panic(err)
	}

	vol, err := strconv.ParseFloat(kline[6], 64)
	if err != nil {
		panic(err)
	}

	quote, err := strconv.ParseFloat(kline[7], 64)
	if err != nil {
		panic(err)
	}
	return &KlineData{
		OpenPrice:  openPrice,
		HighPrice:  highPrice,
		LowsPrice:  lowsPrice,
		ClosePrice: closePrice,
		Vol:        vol,
		Amount:     quote,
		CreatedAt:  createTime / 1000,
	}
}

// InterfaceToStruct 接口转换为结构体
func InterfaceToStruct() {
	data := Message{}
	if err := json.Unmarshal(message, &data); err != nil {
		panic(err)
	}

	tempData := []string{}
	if err := InterfaceToObj(data.Data[0], &tempData); err != nil {
		panic(err)
	}

	fmt.Println("data", tempData)

}

// Ctx 用于管理上下文
type Ctx struct {
	instance context.Context
	close    context.CancelFunc
}

// Ws websocket 实例
type Ws struct {
	ctx         *Ctx            // 上下文实例，管理协程
	instance    *websocket.Conn // 用于关闭协程
	pulse       time.Duration   // 设置脉搏，单位秒（多少秒跳动一次）（多少秒发送一次信息）
	serverAdder string          // 服务器地址
	nor         int             // 重连次数
	Manage      Manage
	onMessage   func(msgType int, data []byte) // 服务获取到的信息，用于给用户处理
	read        func()                         // 读取服务器信息
	heartbeat   func()                         // 检测连接心跳
	close       func()                         // 关闭服务方法
}

type ManageData struct {
	status          int           // 状态：停止，开启，
	isPersistence   bool          // 是否持久化
	PersistenceData []interface{} // 持久化数据
	isCleanUp       bool          // 是否过期清理
	Expiration      time.Duration // 过期时间
}
type Manage interface {
	ManageMessage // 管理信息
	ManageStatus  // 管理状态
}

// ManageStatus 资源管理接口
type ManageStatus interface {
	connect(instance *websocket.Conn)
	status(id string, status int) // 状态管理方法
	read()                        // 读取消息
	heartbeat()                   // 心跳检测
	close(id string)              // 关闭方法
}

// ManageMessage 数据处理方法
type ManageMessage interface {
	OnMessage(msgType int, data []byte)                   // 处理消息
	Persistence(id string, persistenceData []interface{}) // 持久化方法
}

type ManageInstance struct {
}

// status 状态管理方法
func (m *ManageInstance) connect(instance *websocket.Conn) {

}

// status 状态管理方法
func (m *ManageInstance) status(id string, status int) {
}

// read 状态管理方法
func (m *ManageInstance) read() {

}

// heartbeat 状态管理方法
func (m *ManageInstance) heartbeat() {

}

// close 状态管理方法
func (m *ManageInstance) close(id string) {

}

// OnMessage 状态管理方法
func (m *ManageInstance) OnMessage(msgType int, data []byte) {

}

// Persistence 状态管理方法
func (m *ManageInstance) Persistence(id string, persistenceData []interface{}) {

}
