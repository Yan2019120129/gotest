package models

// Level 系统等级配置
type Level struct {
	Id        int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int     `gorm:"type:int unsigned not null;comment:管理ID"`
	Name      string  `gorm:"type:varchar(60) not null;comment:名称"`
	Icon      string  `gorm:"type:varchar(60) not null;comment:图标"`
	Money     float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Days      int     `gorm:"type:tinyint not null;comment:天数"`
	Status    int     `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string  `gorm:"type:text;comment:数据"`
	Desc      string  `gorm:"type:text;comment:详情"`
	UpdatedAt int     `gorm:"type:int unsigned not null;comment:更新时间"`
	CreatedAt int     `gorm:"type:int unsigned not null;comment:创建时间"`
}

const (
	// LevelStatusActive 开启
	LevelStatusActive = 10

	// LevelStatusDisable 禁用
	LevelStatusDisable = -1

	// LevelStatusDelete 删除
	LevelStatusDelete = -2
)
