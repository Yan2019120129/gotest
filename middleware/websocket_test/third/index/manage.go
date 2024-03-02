package index

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	WsStatusStop         = -1
	WsStatusStart        = iota
	WsMessageTypeDefault // 不存储
	WsMessageTypeSub     // 进行持久化，在重连的时候发送到服务器
)

// ManageMessage 数据处理方法
type ManageMessage interface {
	DealWithMessage(msgType int, data []byte)        // 处理消息
	Persistence(msg ...Massage)                      // 持久化方法
	GetPersistence(id string, msgType int) []Massage // 获取持久化数据
}

// Ctx 用于管理上下文
type Ctx struct {
	instance context.Context
	close    context.CancelFunc
}

type Config struct {
	addr, connId string        // 服务地址
	pulse        time.Duration // 脉搏
	nor          int           // 重连次数
}

// Ws websocket 实例
type Ws struct {
	lock          sync.Mutex
	connId        string          // 实例Id
	ctx           *Ctx            // 上下文实例，管理协程
	instance      *websocket.Conn // 用于关闭协程
	pulse         time.Duration   // 设置脉搏，单位秒（多少秒跳动一次）（多少秒发送一次信息）
	serverAdder   string          // 服务器地址
	nor           int             // 重连次数
	ManageMessage                 // 处理信息
}

// NewWs 创建websocket实例
func NewWs(cfg *Config) *Ws {
	if cfg == nil {
		cfg = &Config{
			connId: uuid.NewString(),
			pulse:  5,
			nor:    5,
		}
	}

	// 创建实例
	conn, _, err := websocket.DefaultDialer.Dial(cfg.addr, nil)
	if err != nil {
		logs.Logger.Error(err.Error())
		return nil
	}

	// 设置上下文
	ctx, off := context.WithCancel(context.Background())

	// 配置实例
	ws := &Ws{
		ctx: &Ctx{
			instance: ctx,
			close:    off,
		},
		lock:          sync.Mutex{},
		connId:        cfg.connId,
		instance:      conn,
		pulse:         cfg.pulse,
		serverAdder:   cfg.addr,
		nor:           cfg.nor,
		ManageMessage: &ManageInstance{},
	}
	return ws
}

// close 关闭通道和实例
func (w *Ws) close() {
	logs.Logger.Info("close run")
	defer logs.Logger.Info("close")
	w.ctx.close()
	if err := w.instance.Close(); err != nil {
		logs.Logger.Error(err.Error())
	}
	return
}

// SendMessage 发送消息
func (w *Ws) SendMessage(msg ...Massage) {
	logs.Logger.Info("SendMessageJson", zap.Reflect("msg", msg))
	for _, v := range msg {
		if err := w.instance.WriteMessage(websocket.TextMessage, v.Data); err != nil {
			logs.Logger.Error(err.Error())
		}
	}
	w.Persistence(msg...)
}

// SendMessageJson 发送消息
func (w *Ws) SendMessageJson(msg ...Massage) {
	logs.Logger.Info("SendMessageJson", zap.Reflect("msg", msg))
	for _, v := range msg {
		if err := w.instance.WriteJSON(v.Data); err != nil {
			logs.Logger.Error(err.Error())
		}
	}
	w.Persistence(msg...)
}

// read 读取消息
func (w *Ws) read() {
	logs.Logger.Info("read run")
	defer logs.Logger.Info("read close")
	for {
		select {
		case <-w.ctx.instance.Done():
			return
		default:
			msgType, msg, err := w.instance.ReadMessage()
			if err != nil {
				logs.Logger.Error(err.Error())
				w.connect()
			}
			w.DealWithMessage(msgType, msg)
		}
	}
}

// heartbeat 心脏跳动，固定时间发送请求连接
func (w *Ws) heartbeat() {
	logs.Logger.Info("heartbeat run")
	defer logs.Logger.Info("heartbeat close")
	ch := time.NewTicker(w.pulse * time.Second)
	for {
		select {
		case <-w.ctx.instance.Done():
			return
		default:
			if err := w.instance.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
				logs.Logger.Error(err.Error())
				return
			}
			<-ch.C
		}
	}
}

// connect 重新连接
func (w *Ws) connect() {
	logs.Logger.Info("connect run")
	defer logs.Logger.Info("connect close")
	for i := 0; i < w.nor; i++ {
		var err error
		w.instance, _, err = websocket.DefaultDialer.Dial(w.serverAdder, nil)
		if err != nil {
			logs.Logger.Error("websocket", zap.Int("connect", i), zap.Error(err))
			continue
		}

		// 重连成功发送持久化的订阅数据
		w.SendMessageJson(w.GetPersistence(w.connId, WsMessageTypeSub)...)
		logs.Logger.Info("websocket", zap.Int("connect", i))
		if i == w.nor-1 {
			w.close()
			return
		}
	}
}

// Run 运行 websocket 实例
func (w *Ws) Run() {
	// 发送持久化的订阅数据
	w.SendMessageJson(w.GetPersistence(w.connId, WsMessageTypeSub)...)
	go w.read()
	go w.heartbeat()
}

type ManageInstance struct {
	data []Massage
}

var storagePath = "./data/message.json"

// DealWithMessage 处理消息方法
func (m *ManageInstance) DealWithMessage(msgType int, data []byte) {
	logs.Logger.Info("websocket", zap.Int("type", msgType), zap.String("data", string(data)))
}

// Persistence 数据持久化
func (m *ManageInstance) Persistence(msg ...Massage) {
	logs.Logger.Info("Persistence run")
	// 判断路径是否存在,不存在则创建
	if isPathExist(storagePath) {
		// 获取全部的数据
		if m.data == nil || len(m.data) == 0 {
			m.data = m.GetPersistence("", 0)
		}
		m.data = append(m.data, msg...)
		logs.Logger.Info("Persistence", zap.Reflect("msg", m.data))
		byteData, err := json.Marshal(m.data)
		if err != nil {
			fmt.Println("Marshal err:", err)
			return
		}

		logs.Logger.Info("Persistence", zap.ByteString("msg", byteData))
		// 将数据写入json 文件
		if err = os.WriteFile(storagePath, byteData, 0664); err != nil {
			fmt.Println("WriteFile err:", err)
			return
		}
	}
}

// GetPersistence 获取持久化数据
func (m *ManageInstance) GetPersistence(id string, msgType int) []Massage {
	logs.Logger.Info("GetPersistence run")
	if !isPathExist(storagePath) {
		return nil
	}
	data := make([]Massage, 0)
	if m.data == nil || len(m.data) == 0 {
		// 读取消息
		logs.Logger.Info("GetPersistence", zap.String("path", storagePath))
		storageData, err := os.ReadFile(storagePath)
		if err != nil || storageData == nil || len(storageData) == 0 {
			logs.Logger.Error("read file error:", zap.Error(err))
			return nil
		}
		if err = json.Unmarshal(storageData, &data); err != nil {
			logs.Logger.Error("unmarshal persistence data error:", zap.Error(err))
			return nil
		}
	}

	dataTemp := make([]Massage, 0)
	for _, v := range data {
		switch {
		case id == v.Id && msgType == 0:
			// 获取指定实例全部数据
			dataTemp = append(dataTemp, v)
		case id == v.Id && msgType == v.Type:
			// 获取指定实例，指定类型数据
			dataTemp = append(dataTemp, v)
		default:
			// 获取全部数据
			return data
		}
		return dataTemp
	}
	return nil
}

func isPathExist(path string) bool {
	index := strings.LastIndex(path, "/")
	path = path[:index]
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			logs.Logger.Error("Error creating directory:" + err.Error())
			return false
		}
		logs.Logger.Info("Directory created successfully:" + path)
		return true
	} else if err != nil {
		logs.Logger.Error("Error creating directory:" + err.Error())
		return false
	}
	return true
}
