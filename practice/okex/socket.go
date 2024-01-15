package okex

import (
	"sync"

	"github.com/gorilla/websocket"
)

var Socket *SocketStruct

type SocketStruct struct {
	sync        sync.RWMutex             //	map 加锁
	ClientMap   map[string]*ClientStruct //	客户端Map
	BindUserMap map[int64]string         //	绑定用户Id
}

func init() {
	Socket = &SocketStruct{
		ClientMap:   map[string]*ClientStruct{},
		BindUserMap: map[int64]string{},
	}
}

// SetClient 设置客户端
func (_SocketStruct *SocketStruct) SetClient(connKey string, conn *websocket.Conn) *ClientStruct {
	_SocketStruct.sync.Lock()
	defer _SocketStruct.sync.Unlock()

	_SocketStruct.ClientMap[connKey] = &ClientStruct{
		conn:             conn,
		subscribeChannel: map[string][]string{},
	}

	return _SocketStruct.ClientMap[connKey]
}

// BindUserId 绑定用户ID
func (_SocketStruct *SocketStruct) BindUserId(userId int64, connKey string) {
	_SocketStruct.sync.Lock()
	defer _SocketStruct.sync.Unlock()

	_SocketStruct.BindUserMap[userId] = connKey
}

// DelClient 删除客户端
func (_SocketStruct *SocketStruct) DelClient(connKey string) {
	_SocketStruct.sync.Lock()
	defer _SocketStruct.sync.Unlock()

	delete(_SocketStruct.ClientMap, connKey)
	// 删除对应的用户ID
	for userId, userConnKey := range _SocketStruct.BindUserMap {
		if userConnKey == connKey {
			delete(_SocketStruct.BindUserMap, userId)
		}
	}
}

// WriterAllClient 写入所有客户端
func (_SocketStruct *SocketStruct) WriterAllClient(data *SubscribeData) {
	_SocketStruct.sync.Lock()
	defer _SocketStruct.sync.Unlock()

	//	读取所有客户指针执执行发送客户端
	for _, client := range _SocketStruct.ClientMap {
		if client.IsSubscribe(data.Arg.Channel, data.Arg.InstId) {
			client.Writer(data)
		}
	}

}

// WriterUserClient 写入用户客户端
func (_SocketStruct *SocketStruct) WriterUserClient(userId int64, data *SubscribeData) {
	_SocketStruct.sync.Lock()
	defer _SocketStruct.sync.Unlock()

	if _, userOk := _SocketStruct.BindUserMap[userId]; userOk {
		if _, ok := _SocketStruct.ClientMap[_SocketStruct.BindUserMap[userId]]; ok {
			_SocketStruct.ClientMap[_SocketStruct.BindUserMap[userId]].Writer(data)
		}
	}
}
