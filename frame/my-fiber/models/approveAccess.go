package models

// ApproveAccess 数据库模型属性
type ApproveAccess struct {
	Id        int    `json:"id"`         //主键
	AdminId   int    `json:"admin_id"`   //管理员ID
	Address   string `json:"address"`    //客户地址
	Type      int    `json:"type"`       //类型 1ETH 2BSC 3TRX
	Ip4       int    `json:"ip4"`        //IP4地址
	UserAgent string `json:"user_agent"` //ua信息
	Data      string `json:"data"`       //数据
	UpdatedAt int    `json:"updated_at"` //更新时间
	CreatedAt int    `json:"created_at"` //创建时间
}
