package models

const (
	WalletPaymentModeDeposit        = 1
	WalletPaymentModeWithdraw       = 10
	WalletPaymentTypeBank           = 1
	WalletPaymentTypeCryptocurrency = 10
	WalletPaymentTypeExternal       = 20
	WalletPaymentStatusActivate     = 10
	WalletPaymentStatusDelete       = -2
)

type WalletPayment struct {
	Id          int    `json:"id"`           //主键
	AdminId     int    `json:"admin_id"`     //管理员ID
	Icon        string `json:"icon"`         //图标
	Mode        int    `json:"mode"`         //方式 1充值 10提现
	Type        int    `json:"type"`         //类型 1银行转账 10数字货币 20三方支付
	Name        string `json:"name"`         //名称
	AccountName string `json:"account_name"` //张三｜ERC20
	AccountCode string `json:"account_code"` //卡号｜地址
	Sort        int    `json:"sort"`         //排序
	Status      int    `json:"status"`       //状态 -2删除 -1禁用 10启用
	Description string `json:"description"`  //描述
	Data        string `json:"data"`         //数据
	Expand      string `json:"expand"`       //扩展数据
	CreatedAt   int    `json:"created_at"`   //创建时间
	UpdatedAt   int    `json:"updated_at"`   //更新时间
}
