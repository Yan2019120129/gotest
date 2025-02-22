package client

import "github.com/gomodule/redigo/redis"

// SocketClientInterface socket客户端接口
type SocketClientInterface interface {
	// Connect 连接websocket
	Connect() SocketClientInterface

	// Subscribe 订阅消息
	Subscribe(subscribe *Subscribe) error

	// UnSubscribe 取消订阅
	UnSubscribe(subscribe *Subscribe) error

	// GetSubscribe 获取订阅通道对象
	GetSubscribe(name string) *Subscribe

	// InitSubscribes 初始化订阅数据
	InitSubscribes(subscribeList []*Subscribe) SocketClientInterface

	// ConnWriteMessage 发送消息
	ConnWriteMessage(messageType int, data []byte) error

	// SetWebSocketMessageFunc 设置消息处理
	SetWebSocketMessageFunc(fun func(rdsConn redis.Conn, msg []byte) error) SocketClientInterface

	// SetHeartbeatTime 设置心跳时间
	SetHeartbeatTime(second int) SocketClientInterface

	// SetAddr 设置连接地址
	SetAddr(addr string) SocketClientInterface

	// SetReconnectTime 设置重连时间
	SetReconnectTime(second int) SocketClientInterface

	// SetBeforeConnectFunc 设置重连前置方法
	SetBeforeConnectFunc(func(client SocketClientInterface) error) SocketClientInterface
}
