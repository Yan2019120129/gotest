package okex

import (
	"gotest/practice/okex/utils"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientStruct struct {
	sync             sync.RWMutex        //	map 加锁
	conn             *websocket.Conn     //	客户端对象
	subscribeChannel map[string][]string //	订阅频道
}

// Subscribe 订阅产品
func (_ClientStruct *ClientStruct) Subscribe(msg *Subscribe) {
	_ClientStruct.sync.Lock()
	defer _ClientStruct.sync.Unlock()

	for _, v := range msg.Args {
		//	行情订阅
		if _, ok := _ClientStruct.subscribeChannel[v.Channel]; !ok {
			_ClientStruct.subscribeChannel[v.Channel] = []string{}
		}

		// 新增订阅
		if utils.SliceStringIndexOf(v.InstId, _ClientStruct.subscribeChannel[v.Channel]) == -1 {
			_ClientStruct.subscribeChannel[v.Channel] = append(_ClientStruct.subscribeChannel[v.Channel], v.InstId)
		}
	}
}

// UnSubscribe 取消订阅
func (_ClientStruct *ClientStruct) UnSubscribe(msg *Subscribe) {
	_ClientStruct.sync.Lock()
	defer _ClientStruct.sync.Unlock()

	for _, v := range msg.Args {
		if _, ok := _ClientStruct.subscribeChannel[v.Channel]; ok {
			indexOf := utils.SliceStringIndexOf(v.InstId, _ClientStruct.subscribeChannel[v.Channel])

			if indexOf > -1 {
				_ClientStruct.subscribeChannel[v.Channel] = append(_ClientStruct.subscribeChannel[v.Channel][:indexOf], _ClientStruct.subscribeChannel[v.Channel][indexOf+1:]...)
			}
		}
	}
}

// IsSubscribe 是否订阅
func (_ClientStruct *ClientStruct) IsSubscribe(channel, instId string) bool {
	_ClientStruct.sync.Lock()
	defer _ClientStruct.sync.Unlock()

	if _, ok := _ClientStruct.subscribeChannel[channel]; ok {
		if utils.SliceStringIndexOf(instId, _ClientStruct.subscribeChannel[channel]) > -1 {
			return true
		}
	}
	return false
}

// Writer 发送数据
func (_ClientStruct *ClientStruct) Writer(data any) {
	_ClientStruct.sync.Lock()
	defer _ClientStruct.sync.Unlock()

	if _ClientStruct.conn != nil {
		_ = _ClientStruct.conn.WriteJSON(data)
	}
}
