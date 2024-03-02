package index

import (
	"context"
	"github.com/fasthttp/websocket"
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

// Instance 全局唯一实例
var Instance = &WsInstance{
	WsConnMap: make(map[string]*Ws),
	statusMap: make(map[string]*ManageData),
	lock:      sync.Mutex{},
	Interval:  2,
	once:      sync.Once{},
}

type ManageData struct {
	status        int           // 状态：停止，开启，
	isPersistence bool          // 是否持久化
	isCleanUp     bool          // 是否过期清理
	Expiration    time.Duration // 过期时间
}

// WsInstance websocket 实例
type WsInstance struct {
	once      sync.Once              // 保证状态检测只运行一次
	lock      sync.Mutex             // websocket 实例读写锁
	Interval  time.Duration          // 间隔多少秒检查websocket 状态
	statusMap map[string]*ManageData // 实例状态：启动，关闭
	WsConnMap map[string]*Ws         // 存储 websocket 实例
}

// Massage 发送的消息
type Massage struct {
	Id   string // 实例Id
	Type int    // 数据类型（在启动时会发送订阅类型数据）:订阅类型，默认通知类型
	Data []byte // 数据
}

// NewWs 新建websocket 实例
func (i *WsInstance) NewWs(uuid, addr string) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		logs.Logger.Info("NewWs run")
		// 创建实例
		i.WsConnMap[uuid] = NewWs(&Config{
			addr:   addr,
			connId: uuid,
			pulse:  5,
			nor:    5,
		})

		// 设置状态
		i.statusMap[uuid] = &ManageData{
			status:        WsStatusStop,
			isPersistence: false,
			isCleanUp:     false,
			Expiration:    0,
		}
	}
	i.once.Do(func() {
		i.StatusMonitor()
	})
	return i
}

// Run 运行websocket
func (i *WsInstance) Run(uuids ...string) *WsInstance {
	// 不传参数则启动全部实例
	if len(uuids) < 1 {
		for k, v := range i.WsConnMap {
			if i.statusMap[k].status == WsStatusStart {
				continue
			}
			v.Run()
			i.statusMap[k].status = WsStatusStart
		}
		return i
	}
	// 启动指定实例
	for _, v := range uuids {
		if p, ok := i.WsConnMap[v]; ok {
			if i.statusMap[v].status == WsStatusStart {
				continue
			}
			p.Run()
		} else {
			logs.Logger.Error("could not find " + v)
		}
		logs.Logger.Info("websocket", zap.String(v, "run"))
	}
	return i
}

func (i *WsInstance) StatusMonitor() {
	logs.Logger.Info("StatusMonitor run")
	go func() {
		ch := time.NewTicker(i.Interval * time.Second)
		for {
			for k, v := range i.WsConnMap {
				if err := v.ctx.instance.Err(); err != nil {
					logs.Logger.Error("websocket", zap.String(k, "close"))
				}
				// 状态修改为停止
				i.statusMap[k].status = WsStatusStop
			}
			<-ch.C
		}
	}()
}

// SetWs 设置 websocket。
func (i *WsInstance) SetWs(uuid string, ws *Ws) *WsInstance {
	// 添加协程使用WaitGroup管理线程状态
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid] = ws
		i.statusMap[uuid] = &ManageData{
			status:        WsStatusStop,
			isPersistence: false,
			isCleanUp:     false,
			Expiration:    0,
		}
	}
	return i
}

// ConnectWS 连接okx websocket。
func (i *WsInstance) connect(uuid string) (err error) {
	// 添加协程使用WaitGroup管理线程状态
	if p, ok := i.WsConnMap[uuid]; ok {
		p.instance, _, err = websocket.DefaultDialer.Dial(p.serverAdder, nil)
		if err != nil {
			logs.Logger.Error(err.Error())
			return err
		}
	}
	return nil
}

// SendMessage 发送数据
func (i *WsInstance) SendMessage(msg ...*Massage) *WsInstance {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessage(*v)
		}
	}
	return i
}

// SendMessageJson 发送数据
func (i *WsInstance) SendMessageJson(msg ...*Massage) *WsInstance {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessageJson(*v)
		}
	}
	return i
}

// GetWs 获取ws实例
func (i *WsInstance) GetWs(uuid string) *Ws {
	if _, ok := i.WsConnMap[uuid]; !ok {
		return i.WsConnMap[uuid]
	}
	return nil
}

// setWs 设置ws实例
func (i *WsInstance) setWs(uuid string, ws *Ws) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid] = ws
	}
	return i
}

// SetServerAddr 设置服务器地址
func (i *WsInstance) SetServerAddr(uuid, addr string) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].serverAdder = addr
	}
	return i
}

// SetWsConn 设置服务器地址
func (i *WsInstance) setWsConn(uuid string, conn *websocket.Conn) {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].instance = conn
		i.statusMap[uuid] = &ManageData{
			status:        -1,
			isPersistence: false,
			isCleanUp:     false,
			Expiration:    0,
		}
	}
}

// GetContext 获取上下文
func (i *WsInstance) GetContext(uuid string) *Ctx {
	if _, ok := i.WsConnMap[uuid]; ok {
		return i.WsConnMap[uuid].ctx
	}
	return nil
}

// setContext 设置上下文
func (i *WsInstance) setContext(uuid string) {
	if _, ok := i.WsConnMap[uuid]; ok {
		instance, c := context.WithCancel(context.Background())
		i.WsConnMap[uuid].ctx.instance = instance
		i.WsConnMap[uuid].ctx.close = c
	}
}
