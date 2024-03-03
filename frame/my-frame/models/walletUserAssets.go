package models

// WalletUserAssets 用户钱包资产
type WalletUserAssets struct {
	Id             int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId        int     `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId         int     `gorm:"type:int unsigned not null;comment:用户ID"`
	WalletAssetsId int     `gorm:"type:int unsigned not null;comment:钱包资产ID"`
	Money          float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Status         int     `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data           string  `gorm:"type:text;comment:数据"`
	UpdatedAt      int     `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt      int     `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// WalletUserAssetsStatusActive 开启
	WalletUserAssetsStatusActive = 10

	// WalletUserAssetsStatusDisable 禁用
	WalletUserAssetsStatusDisable = -1

	// WalletUserAssetsStatusDelete 删除
	WalletUserAssetsStatusDelete = -2
)
