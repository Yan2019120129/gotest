package models

// Setting 用户设置
type Setting struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	Name      string `gorm:"type:varchar(50) not null;comment:名称"`
	Type      int    `gorm:"type:tinyint not null default 1;default:1;comment:管理类型一样"`
	Field     string `gorm:"type:varchar(50) not null;uniqueIndex;comment:建铭"`
	Value     string `gorm:"type:text;comment:键值"`
	Data      string `gorm:"type:text;comment:input配置"`
	UpdatedAt int    `gorm:"type:int unsigned not null;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;comment:创建时间"`
}
