package models

// WalletPayment 钱包支付管理
type WalletPayment struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	AssetsId  int    `gorm:"type:int unsigned not null;comment:default:0;钱包资产ID"`
	Name      string `gorm:"type:varchar(60) not null;comment:名称"`
	Icon      string `gorm:"type:varchar(60) not null;comment:图标"`
	Type      int    `gorm:"type:tinyint not null default 1;default:1;comment:类型 1银行卡类型 11数字货币类型 21第三方支付"`
	Mode      int    `gorm:"type:tinyint not null default 1;default:1;comment:模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式"`
	Level     int    `gorm:"type:tinyint not null default 1;default:1;comment:等级"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	Desc      string `gorm:"type:text;comment:详情"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

// WalletPaymentTypeList 支付类型列表
var WalletPaymentTypeList = map[int]string{
	WalletPaymentTypeBank:    "银行卡",
	WalletPaymentTypeDigital: "数字货币",
	WalletPaymentTypeThree:   "第三方支付",
}

// WalletPaymentModeList 支付模式列表
var WalletPaymentModeList = map[int]string{
	WalletPaymentModeDeposit:        "余额充值",
	WalletPaymentModeAssetsDeposit:  "资产充值",
	WalletPaymentModeWithdraw:       "余额提现",
	WalletPaymentModeAssetsWithdraw: "资产提现",
}

const (
	// WalletPaymentTypeBank 银行卡类型
	WalletPaymentTypeBank = 1

	// WalletPaymentTypeDigital 数字货币类型
	WalletPaymentTypeDigital = 11

	// WalletPaymentTypeThree 三方支付
	WalletPaymentTypeThree = 21

	// WalletPaymentModeDeposit 充值模式
	WalletPaymentModeDeposit = 1

	// WalletPaymentModeAssetsDeposit 资产充值模式
	WalletPaymentModeAssetsDeposit = 2

	// WalletPaymentModeWithdraw 提现模式
	WalletPaymentModeWithdraw = 11

	// WalletPaymentModeAssetsWithdraw 资产提现模式
	WalletPaymentModeAssetsWithdraw = 12

	// WalletPaymentStatusActive 开启
	WalletPaymentStatusActive = 10

	// WalletPaymentStatusDisable 禁用
	WalletPaymentStatusDisable = -1

	// WalletPaymentStatusDelete 删除
	WalletPaymentStatusDelete = -2
)
