package models

const (
	ChatConversationTypePrivateLetter = 1  //	私信
	ChatConversationStatusActivate    = 10 //	激活会话
	ChatConversationStatusDelete      = -2 //	删除会话
	ChatConversationStatusShielded    = -1 //	屏蔽会话
)

// ChatConversation 数据库模型属性
type ChatConversation struct {
	Id             int    `json:"id"`              //主键
	ConversationId string `json:"conversation_id"` //会话ID
	UserId         int    `json:"user_id"`         //用户ID
	ReceiverId     int    `json:"receiver_id"`     //接收用户ID
	Type           int    `json:"type"`            //1私聊
	Status         int    `json:"status"`          //状态 -2删除 -1屏蔽 10正常
	Data           string `json:"data"`            //数据
	UpdatedAt      int    `json:"updated_at"`      //更新时间
	CreatedAt      int    `json:"created_at"`      //创建时间
}
