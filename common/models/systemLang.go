package models

// Lang 系统语言
type Lang struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	Name      string `gorm:"type:varchar(60) not null;comment:名称"`
	Alias     string `gorm:"type:varchar(60) not null;comment:别名"`
	Icon      string `gorm:"type:varchar(60) not null;comment:图标"`
	Sort      int    `gorm:"type:tinyint not null;comment:排序"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// LangStatusActive 开启
	LangStatusActive = 10

	// LangStatusDisable 禁用
	LangStatusDisable = -1

	// LangStatusDelete 删除
	LangStatusDelete = -2
)
