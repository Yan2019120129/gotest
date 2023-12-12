package models

// WalletBill 钱包账单
type WalletBill struct {
	Id        int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int     `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int     `gorm:"type:int unsigned not null;comment:用户ID"`
	AssetsId  int     `gorm:"type:int unsigned not null;comment:资产ID"`
	SourceId  int     `gorm:"type:int unsigned not null;comment:来源ID"`
	Type      int     `gorm:"type:smallint not null default 1;default:1;comment:类型 1充值类型 11提现类型 21购买 51收益 61奖励"`
	Name      string  `gorm:"type:varchar(60) not null;comment:名称"`
	Money     float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Balance   float64 `gorm:"type:decimal(12,2) not null;comment:余额"`
	Data      string  `gorm:"type:text;comment:数据"`
	CreatedAt int     `gorm:"type:int unsigned not null;comment:创建时间"`
}

// WalletBillDepositTypeList 入金类型列表
var WalletBillDepositTypeList = map[int]string{
	WalletBillTypeDeposit:              "walletBillTypeDeposit",
	WalletBillTypeWithdrawRefuse:       "walletBillTypeWithdrawRefuse",
	WalletBillTypeAssetsWithdrawRefuse: "walletBillTypeAssetsWithdrawRefuse",
	WalletBillTypeEarnings:             "walletBillTypeEarnings",
	WalletBillTypeAward:                "walletBillTypeAward",
	WalletBillTypeAssetsDeposit:        "walletBillTypeAssetsDeposit",
	WalletBillTypeAssetsEarnings:       "walletBillTypeAssetsEarnings",
	WalletBillTypeAssetsAward:          "walletBillTypeAssetsAward",
	WalletBillTypeRegisterAward:        "walletBillTypeRegisterAward",
	WalletBillTypeShareAward:           "walletBillTypeShareAward",
	WalletBillTypeSystemDeposit:        "walletBillTypeSystemDeposit",
	WalletBillTypeSystemAssetsDeposit:  "walletBillTypeSystemAssetsDeposit",
	WalletBillTypeTeamEarnings:         "walletBillTypeTeamEarnings",
	WalletBillTypeTeamAssetsEarnings:   "walletBillTypeTeamAssetsEarnings",
}

// WalletBillSpendTypeList 花费类型列表
var WalletBillSpendTypeList = map[int]string{
	WalletBillTypeWithdraw:             "walletBillTypeWithdraw",
	WalletBillTypeBuyProduct:           "walletBillTypeBuyProduct",
	WalletBillTypeBuyLevel:             "walletBilTypeBuyLevel",
	WalletBillTypeAssetsWithdraw:       "walletBillTypeAssetsWithdraw",
	WalletBillTypeAssetsBuyProduct:     "walletBillTypeAssetsBuyProduct",
	WalletBillTypeSystemWithdraw:       "walletBillTypeSystemWithdraw",
	WalletBillTypeSystemAssetsWithdraw: "walletBillTypeSystemAssetsWithdraw",
}

const (
	// WalletBillTypeDeposit 充值类型
	WalletBillTypeDeposit = 1

	// WalletBillTypeAssetsDeposit 资产充值类型
	WalletBillTypeAssetsDeposit = 2

	// WalletBillTypeSystemDeposit 余额系统充值
	WalletBillTypeSystemDeposit = 3

	// WalletBillTypeSystemAssetsDeposit 资产系统充值
	WalletBillTypeSystemAssetsDeposit = 4

	// WalletBillTypeWithdraw 提现类型
	WalletBillTypeWithdraw = 11

	// WalletBillTypeAssetsWithdraw 资产提现类型
	WalletBillTypeAssetsWithdraw = 12

	// WalletBillTypeSystemWithdraw 余额系统扣款
	WalletBillTypeSystemWithdraw = 13

	// WalletBilTypeSystemAssetsWithdraw 资产系统扣款
	WalletBillTypeSystemAssetsWithdraw = 14

	// WalletBillTypeWithdrawRefuse 余额提现拒绝
	WalletBillTypeWithdrawRefuse = 15

	// WalletBillTypeAssetsWithdrawRefuse 资产提现拒绝
	WalletBillTypeAssetsWithdrawRefuse = 16

	// WalletBillTypeBuyProduct 购买
	WalletBillTypeBuyProduct = 21

	// WalletBillTypeAssetsBuyProduct 资产购买产品
	WalletBillTypeAssetsBuyProduct = 22

	// WalletBillTypeBuyLevel 购买等级
	WalletBillTypeBuyLevel = 23

	// WalletBillTypeEarnings 收益
	WalletBillTypeEarnings = 51

	// WalletBillTypeAssetsEarnings 资产收益
	WalletBillTypeAssetsEarnings = 52

	// WalletBillTypeAward 奖励
	WalletBillTypeAward = 61

	// WalletBillTypeAssetsAward 资产奖励
	WalletBillTypeAssetsAward = 62

	// WalletBillTypeRegisterAward 注册奖励
	WalletBillTypeRegisterAward = 66

	// WalletBillTypeShareAward 邀请奖励
	WalletBillTypeShareAward = 67

	// WalletBillTypeTeamEarnings 团队收益
	WalletBillTypeTeamEarnings = 71

	// WalletBillTypeTeamAssetsEarnings 团队资产收益
	WalletBillTypeTeamAssetsEarnings = 81
)
