package model

import "encoding/json"

// NetInfo 客户端网络信息
type NetInfo struct {
	Timestamp int64  `json:"timestamp"` // 获取时间戳
	RXBytes   int64  `json:"rx_bytes"`  // 当前速率
	Hostname  string `json:"hostname"`
}

// ReportTcInfo 上报机房带宽
type ReportTcInfo struct {
	KgGroup   string  `json:"kg_group"`
	Bandwidth float64 `json:"bandwidth"`
}

// ResMessage 接口相应信息
type ResMessage struct {
	Code    int             `json:"code"`    // 相应码
	Message string          `json:"message"` // 响应信息
	Data    json.RawMessage `json:"data"`    // 相应数据
}
