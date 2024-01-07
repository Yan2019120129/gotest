package models

const (
	// WalletOrderTypeDeposit 充值
	WalletOrderTypeDeposit = 1
	// WalletOrderTypeSystemDeposit 系统加款
	WalletOrderTypeSystemDeposit = 2
	// WalletOrderTypeWithdraw 提现
	WalletOrderTypeWithdraw = 10
	// WalletOrderTypeSystemWithdraw 系统减款
	WalletOrderTypeSystemWithdraw = 11
	//	WalletOrderTypeAssetsWithdraw 资产提现
	WalletOrderTypeAssetsWithdraw = 12

	WalletOrderStatusPending  = 10
	WalletOrderStatusRefuse   = -1
	WalletOrderStatusComplete = 20
	WalletOrderStatusDelete   = -2
)

type UserWalletOrder struct {
	Id        int     `json:"id"`         //主键
	OrderSn   string  `json:"order_sn"`   //订单号
	AdminId   int     `json:"admin_id"`   //管理员ID
	UserId    int     `json:"user_id"`    //用户ID
	UserType  int     `json:"user_type"`  //用户类型
	Type      int     `json:"type"`       //类型 1充值 2系统加款 10提现 11系统减款
	PaymentId int     `json:"payment_id"` //充值｜提现ID
	Money     float64 `json:"money"`      //金额
	Balance   float64 `json:"balance"`    //余额
	Status    int     `json:"status"`     //状态 -1拒绝 10处理 20完成
	Proof     string  `json:"proof"`      //凭证
	Data      string  `json:"data"`       //数据
	Fee       float64 `json:"fee"`        //手续费
	UpdatedAt int     `json:"updated_at"` //更新时间
	CreatedAt int     `json:"created_at"` //创建时间
}
