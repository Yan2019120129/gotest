package index

import (
	"context"
	"github.com/fasthttp/websocket"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"time"
)

const (
	WsStatusStop  = -1
	WsStatusStart = iota
)

// ManageMessage 数据处理方法
type ManageMessage interface {
	OnMessage(msgType int, data []byte)                                  // 处理消息
	Persistence(id string, persistenceData []interface{}, WsConnMap *Ws) // 持久化方法
}

// Ws websocket 实例
type Ws struct {
	ctx           *Ctx            // 上下文实例，管理协程
	instance      *websocket.Conn // 用于关闭协程
	pulse         time.Duration   // 设置脉搏，单位秒（多少秒跳动一次）（多少秒发送一次信息）
	serverAdder   string          // 服务器地址
	nor           int             // 重连次数
	ManageMessage                 // 处理信息
}

// NewWs 创建websocket实例
func NewWs(addr string, pulse time.Duration, nor int) *Ws {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		logs.Logger.Error(err.Error())
		return nil
	}
	ctx, off := context.WithCancel(context.Background())
	return &Ws{
		ctx: &Ctx{
			instance: ctx,
			close:    off,
		},
		instance:      conn,
		pulse:         pulse,
		serverAdder:   addr,
		nor:           nor,
		ManageMessage: &ManageInstance{},
	}
}

// close 关闭通道和实例
func (w *Ws) close() {
	defer logs.Logger.Info("close ")
	w.ctx.close()
	if err := w.instance.Close(); err != nil {
		logs.Logger.Error(err.Error())
	}
	return
}

// SendMessage 发送消息
func (w *Ws) SendMessage(msg []byte) {
	if err := w.instance.WriteMessage(websocket.TextMessage, msg); err != nil {
		logs.Logger.Error(err.Error())
	}
}

// read 读取消息
func (w *Ws) read() {
	for {
		select {
		case <-w.ctx.instance.Done():
			logs.Logger.Info("read close")
			return
		default:
			msgType, msg, err := w.instance.ReadMessage()
			if err != nil {
				e := err.Error()
				logs.Logger.Error(e)
			}
			w.OnMessage(msgType, msg)
		}
	}
}

// heartbeat 心脏跳动，固定时间发送请求连接
func (w *Ws) heartbeat() {
	defer logs.Logger.Info("heartbeat close ")
	ch := time.NewTicker(w.pulse * time.Second)
	for {
		select {
		case <-w.ctx.instance.Done():
			logs.Logger.Info("heartbeat close ")
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
	logs.Logger.Info("heartbeat run")
	var err error
	w.instance, _, err = websocket.DefaultDialer.Dial(w.serverAdder, nil)
	if err != nil {
		logs.Logger.Error(err.Error())
		panic(err)
	}
}

// Run 运行 websocket 实例
func (w *Ws) Run() {
	go w.read()
	go w.heartbeat()
}

type ManageInstance struct {
}

// OnMessage 放置消息
func (m *ManageInstance) OnMessage(msgType int, data []byte) {
	logs.Logger.Info("websocket", zap.Int("type", msgType), zap.String("data", string(data)))
}

// Persistence 数据持久化
func (m *ManageInstance) Persistence(id string, persistenceData []interface{}, WsConnMap *Ws) {

}
