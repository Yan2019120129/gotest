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
	once:      sync.Once{},
}

// Ctx 用于管理上下文
type Ctx struct {
	instance context.Context
	close    context.CancelFunc
}

type ManageData struct {
	status          int           // 状态：停止，开启，
	isPersistence   bool          // 是否持久化
	PersistenceData []interface{} // 持久化数据
	isCleanUp       bool          // 是否过期清理
	Expiration      time.Duration // 过期时间
}

// WsInstance websocket 实例
type WsInstance struct {
	once      sync.Once
	lock      sync.Mutex             // websocket 实例读写锁
	statusMap map[string]*ManageData // 实例状态：启动，关闭
	WsConnMap map[string]*Ws         // 存储 websocket 实例
}

// NewWs 新建websocket 实例
func (i *WsInstance) NewWs(uuid, addr string) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid] = NewWs(addr, 5, 5)
		i.statusMap[uuid] = &ManageData{
			status:          -1,
			isPersistence:   false,
			PersistenceData: nil,
			isCleanUp:       false,
			Expiration:      0,
		}
	}
	return i
}

// Run 运行websocket
func (i *WsInstance) Run(uuids ...string) *WsInstance {
	// 不传参数则启动全部实例
	if len(uuids) < 1 {
		for k, v := range i.WsConnMap {
			v.Run()
			i.statusMap[k].status = WsStatusStart
		}
		return i
	}
	// 启动指定实例
	for _, v := range uuids {
		if p, ok := i.WsConnMap[v]; ok {
			p.Run()
			i.statusMap[v].status = WsStatusStart
		} else {
			logs.Logger.Error("could not find " + v)
		}
		logs.Logger.Info("websocket", zap.String(v, "run"))
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
func (i *WsInstance) SendMessage(uuid string, msg []byte) *WsInstance {
	i.lock.Lock()
	defer i.lock.Lock()
	if p, ok := i.WsConnMap[uuid]; ok {
		p.SendMessage(msg)
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
