package models

// WalletAssets 钱包资产管理
type WalletAssets struct {
	Id        int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int     `gorm:"type:int unsigned not null;comment:管理ID"`
	Name      string  `gorm:"type:varchar(60) not null;comment:名称"`
	Icon      string  `gorm:"type:varchar(60) not null;comment:图标"`
	Type      int     `gorm:"type:tinyint not null default 1;default:1;comment:类型 1法币资产 11数字货币资产 21虚拟货币资产"`
	Rate      float64 `gorm:"type:decimal(12,2) not null;comment:汇率"`
	Status    int     `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string  `gorm:"type:text;comment:数据"`
	UpdatedAt int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:创建时间"`
}

const (
	// WalletAssetsTypeBank 法币资产
	WalletAssetsTypeFiatCurrency = 1

	// WalletAssetsTypeDigitalCurrency 数字货币资产
	WalletAssetsTypeDigitalCurrency = 11

	// WalletAssetsTypeVirtualCurrency 虚拟货币资产
	WalletAssetsTypeVirtualCurrency = 21

	// WalletAssetsStatusActive 开启
	WalletAssetsStatusActive = 10

	// WalletAssetsStatusDisable 禁用
	WalletAssetsStatusDisable = -1

	// WalletAssetsStatusDelete 删除
	WalletAssetsStatusDelete = -2
)
