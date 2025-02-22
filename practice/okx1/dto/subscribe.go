package dto

import "encoding/json"

// SubscribeParams 发送参数。
type SubscribeParams struct {
	Op   string          `json:"op"`   // 操作，subscribe unsubscribe
	Args []*SubscribeArg `json:"args"` // 请求订阅的频道列表
}

// NewSubscribe 新建订阅
func NewSubscribe() *SubscribeParams {
	return &SubscribeParams{}
}

// Subscribe 设置订阅
func (s *SubscribeParams) Subscribe() *SubscribeParams {
	s.Op = "subscribe"
	return s
}

// Unsubscribe 取消订阅
func (s *SubscribeParams) Unsubscribe() *SubscribeParams {
	s.Op = "unsubscribe"
	return s
}

// SetSubParams 设置订阅订阅参数
func (s *SubscribeParams) SetSubParams(channel, instID string) *SubscribeParams {
	s.Args = append(s.Args, &SubscribeArg{
		Channel: channel,
		InstID:  instID,
	})
	return s
}

// ToString 转换为字符串
func (s *SubscribeParams) ToString() string {
	return string(s.ToBytes())
}

// ToBytes 转换为字节
func (s *SubscribeParams) ToBytes() []byte {
	bytes, _ := json.Marshal(s)
	return bytes
}
