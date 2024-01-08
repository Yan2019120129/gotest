package models

// LevelOrder 用户等级订单
type LevelOrder struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	LevelId   int    `gorm:"type:int unsigned not null;comment:等级ID"`
	Status    int    `gorm:"type:tinyint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	ExpiredAt int    `gorm:"type:int unsigned not null;comment:过期时间"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// UserLevelOrderStatusActive 开启
	UserLevelOrderStatusActive = 10

	// UserLevelOrderStatusDisable 禁用
	UserLevelOrderStatusDisable = -1

	// UserLevelOrderStatusDelete 删除
	UserLevelOrderStatusDelete = -2
)
