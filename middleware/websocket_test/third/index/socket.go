package index

import (
	"context"
	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"sync"
	"time"
)

const (
	MsgTypeErr    = -1
	MsgTypeText   = websocket.TextMessage
	MsgTypeBinary = websocket.BinaryMessage
	MsgTypePong   = websocket.PongMessage
	MsgTypeClose  = websocket.CloseMessage
)

type WsMessageType int

const (
	WsStatusStop                       = -1
	WsStatusStart                      = 1
	WsMessageTypeDefault WsMessageType = iota // 不存储
	WsMessageTypeSub                          // 进行持久化，在重连的时候发送到服务器
)

// Ctx 用于管理上下文
type Ctx struct {
	instance context.Context
	close    context.CancelFunc
}

type Config struct {
	Addr   string        // 服务地址
	ConnId string        // 服务id
	Pulse  time.Duration // 脉搏
	Nor    int           // 重连次数
	ManageMessage
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
	// 默认配置
	if cfg == nil {
		cfg = &Config{
			ConnId:        uuid.NewString(),
			Pulse:         5,
			Nor:           5,
			ManageMessage: &DefaultManage{},
		}
	}

	// 如果没有创建自定义接口则使用默认持久化接口
	if cfg.ManageMessage != nil {
		cfg.ManageMessage = &DefaultManage{}
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
		connId:        cfg.ConnId,
		instance:      getWebSocketInstance(cfg.Addr),
		pulse:         cfg.Pulse,
		serverAdder:   cfg.Addr,
		nor:           cfg.Nor,
		ManageMessage: cfg.ManageMessage,
	}

	return ws
}

// send 用于发送数据
func (w *Ws) send(msg []byte) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	if err := w.instance.WriteMessage(websocket.TextMessage, msg); err != nil {
		logs.Logger.Error(err.Error())
		return err
	}
	return nil
}

// SendMessage 发送消息
func (w *Ws) SendMessage(msg ...Massage) {
	logs.Logger.Info("SendMessageJson", zap.Reflect("msg", msg))
	for _, v := range msg {
		if err := w.send(v.Data); err != nil {
			logs.Logger.Error(err.Error())
			return
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
				w.reconnection()
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
			if err := w.send([]byte("ping")); err != nil {
				logs.Logger.Error(err.Error())
				return
			}
			<-ch.C
		}
	}
}

// reconnection 重新连接
func (w *Ws) reconnection() {
	logs.Logger.Info("reconnection run")
	defer logs.Logger.Info("reconnection close")
	for i := 0; i < w.nor; i++ {
		var err error
		w.instance = getWebSocketInstance(w.serverAdder)
		if err != nil {
			logs.Logger.Error("websocket", zap.Int("reconnection", i), zap.Error(err))
			continue
		}

		logs.Logger.Info("websocket", zap.Int("reconnection", i))
		if i == w.nor-1 {
			w.close()
			return
		}
	}
}

// rsSub 重新订阅
func (w *Ws) resubscribe() {
	// 重连成功发送持久化的订阅数据
	logs.Logger.Info("resubscribe", zap.String("id", w.connId))
	for _, v := range w.GetPersistence(w.connId, WsMessageTypeSub) {
		logs.Logger.Info("resubscribe", zap.String("id", v.Id))
		if err := w.send(v.Data); err != nil {
			logs.Logger.Error("resubscribe error", zap.Error(err))
			continue
		}
	}
}

// Run 运行 websocket 实例
func (w *Ws) Run() {
	// 发送持久化的订阅数据
	w.resubscribe()
	go w.read()
	go w.heartbeat()
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

// getWebSocketInstance 获取websocket实例
func getWebSocketInstance(addr string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		logs.Logger.Error("getWebSocketInstance err", zap.Error(err))
		return nil

	}
	return conn
}
