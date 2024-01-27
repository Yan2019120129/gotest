package models

// WalletOrder 钱包订单
type WalletOrder struct {
	Id        int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int     `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int     `gorm:"type:int unsigned not null;comment:用户ID"`
	AssetsId  int     `gorm:"type int unsigned not null;comment:资产ID"`
	SourceId  int     `gorm:"type:int unsigned not null;comment:来源ID"`
	Type      int     `gorm:"type:tinyint not null default 1;default:1;comment:类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型"`
	OrderSn   string  `gorm:"type:varchar(60) not null;uniqueIndex;comment:编号"`
	Money     float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Fee       float64 `gorm:"type:decimal(5,2) not null;comment:手续费"`
	Voucher   string  `gorm:"type:varchar(250) not null;comment:支付凭证"`
	Status    int     `gorm:"type:smallint not null default 10;default:10;comment:状态  -2删除 -1禁用 10审核 20完成"`
	Data      string  `gorm:"type:text;comment:数据"`
	UpdatedAt int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int     `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// WalletOrderTypeDeposit 充值类型
	WalletOrderTypeDeposit = 1

	// WalletOrderTypeAssetsDeposit 资产充值类型
	WalletOrderTypeAssetsDeposit = 2

	// WalletOrderTypeWithdraw 提现类型
	WalletOrderTypeWithdraw = 11

	// WalletOrderTypeAssetsWithdraw 资产提现类型
	WalletOrderTypeAssetsWithdraw = 12

	// WalletOrderStatusComplete 完成
	WalletOrderStatusComplete = 20

	// WalletOrderStatusActive 审核
	WalletOrderStatusActive = 10

	// WalletOrderStatusRefuse 拒绝
	WalletOrderStatusRefuse = -1

	// WalletOrderStatusDelete 删除
	WalletOrderStatusDelete = -2
)
