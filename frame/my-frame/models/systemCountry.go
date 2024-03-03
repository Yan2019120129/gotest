package models

// Country 系统国家
type Country struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	Name      string `gorm:"type:varchar(60) not null;comment:名称"`
	Alias     string `gorm:"type:varchar(60) not null;comment:别名"`
	Icon      string `gorm:"type:varchar(60) not null;comment:图标"`
	Iso1      string `gorm:"type:varchar(60) not null;comment:二位代码"`
	Sort      int    `gorm:"type:tinyint not null;default:99;comment:排序"`
	Code      string `gorm:"type:varchar(50) not null;comment:区号"`
	Status    int    `gorm:"type:smallint not null default 10;default:10;comment:状态 -2删除 -1禁用 10开启"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// CountryStatusActive 开启
	CountryStatusActive = 10

	// CountryStatusDisable 禁用
	CountryStatusDisable = -1

	// CountryStatusDelete 删除
	CountryStatusDelete = -2
)
