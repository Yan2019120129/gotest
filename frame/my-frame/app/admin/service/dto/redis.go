package dto

// RdsPublishAndParams redis发布订阅参数
type RdsPublishAndParams struct {
	Channel string `json:"channel"` // 订阅通道
	Message string `json:"message"` // 内容
	Type    int    `json:"type"`    // 类型 1发布 2订阅
}

// RdsParams 设置redis参数
type RdsParams struct {
	Command string `json:"command"` // 命令
	Type    int    `json:"type"`    // 类型 1添加 2获取
	Key     string `json:"key"`     // 键
	Value   string `json:"value"`   // 值
}
