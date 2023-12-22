package models

// ApproveAddressIncome 数据库模型属性
type ApproveAddressIncome struct {
	Id        int     `json:"id"`         //主键
	AdminId   int     `json:"admin_id"`   //管理员ID
	AddressId int     `json:"address_id"` //授权地址ID
	Type      int     `json:"type"`       //类型
	Balance   float64 `json:"balance"`    //余额
	Money     float64 `json:"money"`      //金额
	Data      string  `json:"data"`       //数据
	CreatedAt int     `json:"created_at"` //创建时间
}

const (
	ApproveAddressIncomeAuto int = 20 //		自动收益
)
