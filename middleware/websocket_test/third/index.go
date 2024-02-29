package third

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
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

var Instance = &WsInstance{
	WsConnMap: make(map[string]*Ws),
}

type Ws struct {
	ctx          context.Context
	instance     *websocket.Conn
	serverAdder  string
	pulse        int // 设置脉搏，单位秒（多少秒跳动一次）（多少秒发送一次信息）
	onMessageFun func(msgType int, data []byte)
	readFun      func()
	heartbeatFun func()
	closeFun     func()
}

// WsInstance websocket 实例
type WsInstance struct {
	lock      sync.Mutex
	WsConnMap map[string]*Ws
}

// NewWs 新建websocket 实例
func (i *WsInstance) NewWs(uuid, addr string) *WsInstance {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		logs.Logger.Error(err.Error())
		return nil
	}

	// 生成UUID 创建实例，并返回UUID给创建者
	return i.setWs(uuid, &Ws{ctx: context.Background(), instance: conn, serverAdder: addr, pulse: 5}).
		setDefaultReadFunc(uuid).
		setDefaultHeartbeat(uuid).
		SetDefaultOnMessageFunc(uuid)
}

// Run 运行websocket
func (i *WsInstance) Run() {
	for _, v := range i.WsConnMap {
		go v.readFun()
		go v.heartbeatFun()
	}
}

// SendMessage 发送数据
func (i *WsInstance) SendMessage(uuid string, msg []byte) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		if err := p.instance.WriteMessage(websocket.TextMessage, msg); err != nil {
			logs.Logger.Error(err.Error())
		}
	}
	return i
}

func (i *WsInstance) SetCloseFunc(uuid string) {
	if p, ok := i.WsConnMap[uuid]; ok {
		closeFun := func() {
			defer logs.Logger.Info("close ")
			_, off := context.WithCancel(p.ctx)
			if err := p.instance.Close(); err != nil {
				logs.Logger.Error(err.Error())
			}
			off()
		}
		p.closeFun = closeFun
	}
}

// setDefaultHeartbeat 设置心跳方法
func (i *WsInstance) setDefaultHeartbeat(uuid string) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		if p.ctx == nil {
			p.ctx = context.Background()
		}
		//done := p.ctx.Done()
		heartFun := func() {
			defer logs.Logger.Info("heartbeat close ")
			ch := time.NewTicker(time.Duration(p.pulse) * time.Second)
			for {
				select {
				case <-p.ctx.Done():
					return
				default:
					if err := p.instance.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
						logs.Logger.Error(err.Error())
						return
					}
					<-ch.C
				}
			}
		}
		p.heartbeatFun = heartFun
	}
	return i
}

// GetDefaultOnMessageFunc 获取默认的信息放置
func (i *WsInstance) GetDefaultOnMessageFunc() func(msgType int, data []byte) {
	return func(msgType int, data []byte) {
		fmt.Println("type：", msgType, "data：", string(data))
	}
}

// SetDefaultOnMessageFunc 设置默认的信息放置
func (i *WsInstance) SetDefaultOnMessageFunc(uuid string) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid].onMessageFun = i.GetDefaultOnMessageFunc()
	}
	return i
}

// setDefaultReadFunc 设置默认的读取方法
func (i *WsInstance) setDefaultReadFunc(uuid string) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		//done := p.ctx.Done()
		readFun := func() {
			defer logs.Logger.Info("read close")
			I := 0
			for {
				select {
				case <-p.ctx.Done():
					return
				default:
					msgType, msg, err := p.instance.ReadMessage()
					if err != nil {
						e := err.Error()
						logs.Logger.Error(e)
						//p.onMessageFun(msgType, []byte(e))
						p.closeFun()
					}
					p.onMessageFun(msgType, msg)
				}
				I++
				fmt.Println(I)
				if I == 20 {
					p.closeFun()
				}
			}
		}
		p.readFun = readFun
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
func (i *WsInstance) GetContext(uuid string) *context.Context {
	if _, ok := i.WsConnMap[uuid]; ok {
		return &i.WsConnMap[uuid].ctx
	}
	return nil
}

// setContext 设置上下文
func (i *WsInstance) setContext(uuid string, ctx context.Context) {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].ctx = ctx
	}
}

// SetOnMessageFun 设置消息处理方法
func (i *WsInstance) SetOnMessageFun(uuid string, fu func(msgType int, data []byte)) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].onMessageFun = fu
	}
	return i
}

// SetReadFun 设置读取消息处理方法
func (i *WsInstance) SetReadFun(uuid string, fu func()) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].readFun = fu
	}
	return i
}

// SetCloseFun 设置读取消息处理方法
func (i *WsInstance) SetCloseFun(uuid string, fu func()) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].closeFun = fu
	}
	return i
}
