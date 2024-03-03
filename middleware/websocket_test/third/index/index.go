package index

import (
	"context"
	"github.com/fasthttp/websocket"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"sync"
	"time"
)

// Instance 全局唯一实例
var Instance = &WsManage{
	WsConnMap: make(map[string]*Ws),
	statusMap: make(map[string]*WsStatus),
	lock:      sync.Mutex{},
	Interval:  2,
	once:      sync.Once{},
}

// WsStatus websocket 实例状态处理
type WsStatus struct {
	status        int           // 状态：停止，开启，
	isPersistence bool          // 是否持久化
	isCleanUp     bool          // 是否过期清理
	Expiration    time.Duration // 过期时间
}

// WsManage websocket 实例
type WsManage struct {
	once      sync.Once            // 保证状态检测只运行一次
	lock      sync.Mutex           // websocket 实例读写锁
	Interval  time.Duration        // 间隔多少秒检查websocket 状态
	statusMap map[string]*WsStatus // 实例状态：启动，关闭
	WsConnMap map[string]*Ws       // 存储 websocket 实例
}

// NewWsManage 新建websocket 管理实例
func NewWsManage(Interval time.Duration) *WsManage {
	return &WsManage{
		once:      sync.Once{},
		lock:      sync.Mutex{},
		Interval:  Interval,
		statusMap: make(map[string]*WsStatus),
		WsConnMap: make(map[string]*Ws),
	}
}

// NewWs 新建websocket 实例
func (i *WsManage) NewWs(uuid, addr string) *WsManage {
	if _, ok := i.WsConnMap[uuid]; !ok {
		logs.Logger.Info("NewWs run")
		// 创建实例
		i.WsConnMap[uuid] = NewWs(&Config{
			Addr:   addr,
			ConnId: uuid,
			Pulse:  5,
			Nor:    5,
		})

		// 设置状态
		i.statusMap[uuid] = &WsStatus{
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
func (i *WsManage) Run(uuids ...string) *WsManage {
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

func (i *WsManage) StatusMonitor() {
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
func (i *WsManage) SetWs(uuid string, ws *Ws) *WsManage {
	// 添加协程使用WaitGroup管理线程状态
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid] = ws
		i.statusMap[uuid] = &WsStatus{
			status:        WsStatusStop,
			isPersistence: false,
			isCleanUp:     false,
			Expiration:    0,
		}
	}
	return i
}

// GetUUIDAll 获取全部实例的UUID
func (i *WsManage) GetUUIDAll() []string {
	i.lock.Lock()
	defer i.lock.Unlock()
	uuids := make([]string, 0)
	for k, _ := range i.WsConnMap {
		uuids = append(uuids, k)
	}
	return uuids
}

// ConnectWS 连接okx websocket。
func (i *WsManage) reconnection(uuids ...string) (err error) {
	for _, v := range uuids {
		if p, ok := i.statusMap[v]; ok {
			switch p.status {
			case WsStatusStart:
			case WsStatusStop:
				i.setWsConn(v, getWebSocketInstance(v))
				if err != nil {
					logs.Logger.Error(err.Error())
					return err
				}
			default:
				panic("unhandled default case")
			}

		}
	}
	return nil
}

// SendMessage 发送数据
func (i *WsManage) SendMessage(msg ...*Massage) *WsManage {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessage(*v)
		}
	}
	return i
}

// SendMessageJson 发送数据
func (i *WsManage) SendMessageJson(msg ...*Massage) *WsManage {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessageJson(*v)
		}
	}
	return i
}

// GetWs 获取ws实例
func (i *WsManage) GetWs(uuid string) *Ws {
	if _, ok := i.WsConnMap[uuid]; !ok {
		return i.WsConnMap[uuid]
	}
	return nil
}

// setWs 设置ws实例
func (i *WsManage) setWs(uuid string, ws *Ws) *WsManage {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid] = ws
	}
	return i
}

// GetServerAddr 获取服务地址
func (i *WsManage) GetServerAddr(uuid string) string {
	if p, ok := i.WsConnMap[uuid]; !ok {
		return p.serverAdder
	}
	return ""
}

// SetServerAddr 设置服务器地址
func (i *WsManage) SetServerAddr(uuid, addr string) *WsManage {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].serverAdder = addr
	}
	return i
}

// GetWsConn 获取websocket实例
func (i *WsManage) GetWsConn(uuid string) *websocket.Conn {
	if p, ok := i.WsConnMap[uuid]; !ok {
		return p.instance
	}
	return nil
}

// SetWsConn 设置服务器地址
func (i *WsManage) setWsConn(uuid string, conn *websocket.Conn) {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].instance = conn
		i.statusMap[uuid] = &WsStatus{
			status:        -1,
			isPersistence: false,
			isCleanUp:     false,
			Expiration:    0,
		}
	}
}

// GetContext 获取上下文
func (i *WsManage) GetContext(uuid string) *Ctx {
	if _, ok := i.WsConnMap[uuid]; ok {
		return i.WsConnMap[uuid].ctx
	}
	return nil
}

// setContext 设置上下文
func (i *WsManage) setContext(uuid string) {
	if _, ok := i.WsConnMap[uuid]; ok {
		instance, c := context.WithCancel(context.Background())
		i.WsConnMap[uuid].ctx.instance = instance
		i.WsConnMap[uuid].ctx.close = c
	}
}
