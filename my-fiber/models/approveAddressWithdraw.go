package models

const (
	ApproveAddressWithdrawStatusPending = 10
	ApproveAddressWithdrawStatusRefuse  = -1
	ApproveAddressWithdrawStatusDelete  = -2
)

// ApproveAddressWithdraw 数据库模型属性
type ApproveAddressWithdraw struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	AddressId int     `json:"address_id"` //授权地址ID
	Type      int     `json:"type"`       //类型
	Balance   float64 `json:"balance"`    //余额
	Money     float64 `json:"money"`      //金额
	Status    int     `json:"status"`     //状态 -2删除 -1拒绝 10处理 20完成
	Data      string  `json:"data"`       //数据
	UpdatedAt int     `json:"updated_at"` //更新时间
	CreatedAt int     `json:"created_at"` //创建时间
}
