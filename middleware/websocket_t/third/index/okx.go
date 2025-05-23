package index

import (
	"encoding/json"
	"gotest/common/module/logs"

	"go.uber.org/zap"
)

const uuidOkx = "okx"

// OkxParams 发送参数。
type OkxParams struct {
	Op   string `json:"op"`   // 操作，subscribe unsubscribe
	Args []*Arg `json:"args"` // 请求订阅的频道列表
}

// Arg 币种订阅频道。
type Arg struct {
	Channel string `json:"channel"` // 订阅的通道
	InstID  string `json:"instId"`  // 货币类型
}

// OkxManage 管理数据管理
type OkxManage struct {
	data []Massage
}

// DealWithMessage 处理消息方法
func (m *OkxManage) DealWithMessage(msgType int, data []byte) {
	logs.Logger.Info("websocket", zap.Int("type", msgType), zap.String("data", string(data)))
}

// Persistence 数据持久化
func (m *OkxManage) Persistence(msg ...Massage) {

}

// GetPersistence 获取持久化数据
func (m *OkxManage) GetPersistence(id string, msgType WsMessageType) []Massage {
	if id == uuidOkx && msgType == WsMessageTypeSub {
		tempData := OkxParams{
			Op:   "subscribe",
			Args: []*Arg{},
		}
		var instIds []string
		for _, v := range instIds {
			tempData.Args = append(tempData.Args, []*Arg{
				{Channel: ChannelTicker, InstID: v},
				{Channel: ChannelTrades, InstID: v},
				{Channel: ChannelBooks, InstID: v},
				{Channel: ChannelBooks5, InstID: v},
			}...)
		}
		okxSubData, err := json.Marshal(tempData)
		if err != nil {
			logs.Logger.Error(logs.LogMsgOkx, zap.String("method", "getInstIds"), zap.Error(err))
			return nil
		}
		return []Massage{
			{Id: uuidOkx, Type: WsMessageTypeSub, Data: okxSubData},
		}
	}
	return nil
}
