package client

import (
	"gotest/common/module/cache"
	"sync"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

// SocketClient socket客户端
type SocketClient struct {
	sync                 sync.Mutex                                 //	锁
	err                  error                                      //	错误信息
	addr                 string                                     // 	连接地址
	conn                 *websocket.Conn                            //	连接对象
	subscribes           []*Subscribe                               //	订阅数据
	heartbeatTime        time.Duration                              //	心跳时间
	reconnectTime        time.Duration                              //	重连时间
	beforeConnectFunc    func(client SocketClientInterface) error   //	连接前执行
	websocketMessageFunc func(rdsConn redis.Conn, msg []byte) error //	websocket消息处理方法
}

// NewSocketClient 创建新的连接对象
func NewSocketClient(addr string) SocketClientInterface {
	client := &SocketClient{
		addr:          addr,
		sync:          sync.Mutex{},
		heartbeatTime: 10 * time.Second,
		reconnectTime: 10 * time.Second,
		subscribes:    make([]*Subscribe, 0),
		websocketMessageFunc: func(rdsConn redis.Conn, msg []byte) error {
			return nil
		},
	}
	return client
}

// Run 运行读取
func (_SocketClient *SocketClient) run() error {
	if _SocketClient.err != nil {
		return _SocketClient.err
	}

	// 订阅数据
	for _, subscribe := range _SocketClient.subscribes {
		if len(subscribe.Data) == 0 {
			continue
		}
		_ = _SocketClient.ConnWriteMessage(websocket.TextMessage, subscribe.Data)
	}

	go func() {
		rdsCon := cache.RdsPool.Get()
		defer func(rdsCon redis.Conn) {
			// 重新连接
			_SocketClient.reconnect()
			_ = rdsCon.Close()
		}(rdsCon)

		var msg []byte
		var err error
		for {
			_, msg, err = _SocketClient.conn.ReadMessage()
			if err != nil {
				return
			}

			// 读取消息
			_ = _SocketClient.websocketMessageFunc(rdsCon, msg)
		}
	}()
	return nil
}

// SetWebSocketMessageFunc 设置消息处理
func (_SocketClient *SocketClient) SetWebSocketMessageFunc(fun func(rdsConn redis.Conn, msg []byte) error) SocketClientInterface {
	_SocketClient.websocketMessageFunc = fun
	return _SocketClient
}

// InitSubscribes 初始化订阅数据
func (_SocketClient *SocketClient) InitSubscribes(subscribeList []*Subscribe) SocketClientInterface {
	_SocketClient.subscribes = subscribeList
	return _SocketClient
}

// Subscribe 订阅消息
func (_SocketClient *SocketClient) Subscribe(subscribe *Subscribe) error {
	oldSubscribe := _SocketClient.GetSubscribe(subscribe.Name)
	if oldSubscribe == nil {
		_SocketClient.subscribes = append(_SocketClient.subscribes, subscribe)
	}
	dataBytes, _ := json.Marshal(subscribe.Data)
	return _SocketClient.ConnWriteMessage(websocket.TextMessage, dataBytes)
}

// UnSubscribe 取消订阅
func (_SocketClient *SocketClient) UnSubscribe(subscribe *Subscribe) error {
	for i := 0; i < len(_SocketClient.subscribes); i++ {
		if _SocketClient.subscribes[i].Name == subscribe.Name {
			_SocketClient.subscribes = append(_SocketClient.subscribes[:i], _SocketClient.subscribes[i+1:]...)
			dataBytes, _ := json.Marshal(subscribe.Data)
			return _SocketClient.ConnWriteMessage(websocket.TextMessage, dataBytes)
		}
	}
	return nil
}

// GetSubscribe 获取订阅通道对象
func (_SocketClient *SocketClient) GetSubscribe(name string) *Subscribe {
	for _, subscribe := range _SocketClient.subscribes {
		if subscribe.Name == name {
			return subscribe
		}
	}
	return nil
}

// ConnWriteMessage 发送字节消息
func (_SocketClient *SocketClient) ConnWriteMessage(messageType int, data []byte) error {
	_SocketClient.sync.Lock()
	defer _SocketClient.sync.Unlock()

	return _SocketClient.conn.WriteMessage(messageType, data)
}

// SetHeartbeatTime 设置心跳时间
func (_SocketClient *SocketClient) SetHeartbeatTime(second int) SocketClientInterface {
	_SocketClient.heartbeatTime = time.Duration(second) * time.Second
	return _SocketClient
}

// SetAddr 设置连接地址
func (_SocketClient *SocketClient) SetAddr(addr string) SocketClientInterface {
	_SocketClient.addr = addr
	return _SocketClient
}

// SetReconnectTime 设置重连时间
func (_SocketClient *SocketClient) SetReconnectTime(second int) SocketClientInterface {
	_SocketClient.reconnectTime = time.Duration(second) * time.Second
	return _SocketClient
}

// SetBeforeConnectFunc 设置连接前执行
func (_SocketClient *SocketClient) SetBeforeConnectFunc(fun func(client SocketClientInterface) error) SocketClientInterface {
	_SocketClient.beforeConnectFunc = fun
	return _SocketClient
}

// Connect 连接
func (_SocketClient *SocketClient) Connect() SocketClientInterface {
	dialer := &websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
	}
	_SocketClient.conn, _, _SocketClient.err = dialer.Dial(_SocketClient.addr, nil)

	//	如果连接成功启动心跳处理
	if _SocketClient.err == nil && _SocketClient.heartbeatTime > 0 {
		go func() {
			ch := time.NewTicker(_SocketClient.heartbeatTime)

			for {
				if err := _SocketClient.ConnWriteMessage(websocket.PingMessage, []byte{}); err != nil {
					_ = _SocketClient.conn.Close()
					return
				}
				<-ch.C
			}
		}()
	}

	// 启动读取
	_ = _SocketClient.run()
	return _SocketClient
}

// Reconnect 重新连接
func (_SocketClient *SocketClient) reconnect() {
	time.AfterFunc(_SocketClient.reconnectTime, func() {
		// 如果设置了重连前置方法, 那么判断是否正确
		if _SocketClient.beforeConnectFunc != nil {
			err := _SocketClient.beforeConnectFunc(_SocketClient)
			if err != nil {
				_SocketClient.reconnect()
				return
			}
		}

		_SocketClient.Connect()
		if _SocketClient.err != nil {
			_SocketClient.reconnect()
			return
		}
	})
}
