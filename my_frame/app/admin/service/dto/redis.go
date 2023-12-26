package dto

// RdsPublishAndParams redis发布订阅参数
type RdsPublishAndParams struct {
	Channel string `json:"channel"` // 订阅通道
	Message string `json:"message"` // 内容
	Type    string `json:"type"`    // 类型 1发布 2订阅
}
