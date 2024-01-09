package models

// WalletUserAccount 钱包卡片管理
type WalletUserAccount struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	PaymentId int    `gorm:"type:int unsigned not null;comment:支付ID"`
	Name      string `gorm:"type:varchar(50) not null;comment:银行名称｜Token"`
	RealName  string `gorm:"type:varchar(50) not null;comment:真实姓名"`
	Number    string `gorm:"type:varchar(50) not null;comment:卡号|地址"`
	Code      string `gorm:"type:varchar(255) not null;comment:银行代码"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// WalletUserAccountStatusActive 开启
	WalletUserAccountStatusActive = 10

	// WalletUserAccountStatusDisable 禁用
	WalletUserAccountStatusDisable = -1

	// WalletUserAccountStatusDelete 删除
	WalletUserAccountStatusDelete = -2
)
