package models

const (
	ChatConversationMessageUnread = 1 //	消息未读
	ChatConversationMessageRead   = 2 //	消息已读
	ChatConversationTypeText      = 1 //	文本
	ChatConversationTypeImage     = 2 //	图片
)

// ChatConversationMessage 数据库模型属性
type ChatConversationMessage struct {
	Id             int    `json:"id"`              //主键
	ConversationId string `json:"conversation_id"` //会话ID
	SenderId       int    `json:"sender_id"`       //发送者ID
	ReceiverId     int    `json:"receiver_id"`     //接收者ID
	Unread         int    `json:"unread"`          //未读 1未读 2已读
	Type           int    `json:"type"`            //消息类型 1文本 2图片 3语音 4视频 10富文本
	Data           string `json:"data"`            //消息内容
	Extra          string `json:"extra"`           //扩展数据
	UpdatedAt      int    `json:"updated_at"`      //更新时间
	CreatedAt      int    `json:"created_at"`      //创建时间
}
