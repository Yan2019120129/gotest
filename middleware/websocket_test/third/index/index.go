package index

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

type Ctx struct {
	instance context.Context
	close    context.CancelFunc
}

type Ws struct {
	ctxInstance *Ctx                           // 上下文实例，管理协程
	instance    *websocket.Conn                // 用于关闭协程
	serverAdder string                         // 服务器地址
	pulse       time.Duration                  // 设置脉搏，单位秒（多少秒跳动一次）（多少秒发送一次信息）
	nor         int                            // 重连次数
	onMessage   func(msgType int, data []byte) // 服务获取到的信息，用于给用户处理
	read        func()                         // 读取服务器信息
	heartbeat   func()                         // 检测连接心跳
	close       func()                         // 关闭服务方法
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
	ins, off := context.WithCancel(context.Background())
	return i.setWs(uuid, &Ws{
		ctxInstance: &Ctx{
			instance: ins,
			close:    off,
		}, instance: conn, serverAdder: addr, pulse: 5}).
		setDefaultReadFunc(uuid).
		setDefaultHeartbeatFun(uuid).
		SetDefaultOnMessageFunc(uuid).
		SetCloseFunc(uuid)
}

// Run 运行websocket
func (i *WsInstance) Run() {
	for _, v := range i.WsConnMap {
		go v.read()
		go v.heartbeat()
	}
}

// SendMessage 发送数据
func (i *WsInstance) SendMessage(uuid string, msg []byte) *WsInstance {
	i.lock.Lock()
	defer i.lock.Lock()
	if p, ok := i.WsConnMap[uuid]; ok {
		if err := p.instance.WriteMessage(websocket.TextMessage, msg); err != nil {
			logs.Logger.Error(err.Error())
		}
	}
	return i
}

func (i *WsInstance) SetCloseFunc(uuid string) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		closeFun := func() {
			defer logs.Logger.Info("close ")
			p.ctxInstance.close()
			if err := p.instance.Close(); err != nil {
				logs.Logger.Error(err.Error())
			}
			return
		}
		p.close = closeFun
	}
	return i
}

// setDefaultHeartbeatFun 设置心跳方法
func (i *WsInstance) setDefaultHeartbeatFun(uuid string) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		heartFun := func() {
			//defer logs.Logger.Info("heartbeat close ")
			ch := time.NewTicker(time.Duration(p.pulse) * time.Second)
			for {
				select {
				case <-p.ctxInstance.instance.Done():
					logs.Logger.Info("heartbeat close ")
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
		p.heartbeat = heartFun
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
		i.WsConnMap[uuid].onMessage = i.GetDefaultOnMessageFunc()
	}
	return i
}

// setDefaultReadFunc 设置默认的读取方法
func (i *WsInstance) setDefaultReadFunc(uuid string) *WsInstance {
	if p, ok := i.WsConnMap[uuid]; ok {
		readFun := func() {
			I := 0
			for {
				select {
				case <-p.ctxInstance.instance.Done():
					logs.Logger.Info("read close")
					return
				default:
					msgType, msg, err := p.instance.ReadMessage()
					if err != nil {
						e := err.Error()
						logs.Logger.Error(e)
						//p.onMessage(msgType, []byte(e))
						//p.close()
					}
					p.onMessage(msgType, msg)
				}
				I++
				fmt.Println(I)
				if I == 10 {
					p.close()
				}
			}
		}
		p.read = readFun
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
		return i.WsConnMap[uuid].ctxInstance
	}
	return nil
}

// setContext 设置上下文
func (i *WsInstance) setContext(uuid string) {
	if _, ok := i.WsConnMap[uuid]; ok {
		instance, c := context.WithCancel(context.Background())
		i.WsConnMap[uuid].ctxInstance.instance = instance
		i.WsConnMap[uuid].ctxInstance.close = c
	}
}

// SetOnMessageFun 设置消息处理方法
func (i *WsInstance) SetOnMessageFun(uuid string, fu func(msgType int, data []byte)) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid].onMessage = fu
	}
	return i
}

// SetReadFun 设置读取消息处理方法
func (i *WsInstance) SetReadFun(uuid string, fu func()) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid].read = fu
	}
	return i
}

// SetCloseFun 设置读取消息处理方法
func (i *WsInstance) SetCloseFun(uuid string, fu func()) *WsInstance {
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid].close = fu
	}
	return i
}
