package models

// Access 前台访问表
type Access struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int    `gorm:"type:int unsigned not null;comment:用户ID"`
	Name      string `gorm:"type:varchar(120) not null;comment:名称"`
	Ip        int    `gorm:"type:int unsigned not null;comment:IP4地址"`
	Type      int    `gorm:"type:tinyint not null default 1;default:1;comment:类型 1访问量"`
	Headers   string `gorm:"type:text;comment:头信息"`
	Data      string `gorm:"type:text;comment:数据"`
	UpdatedAt int    `gorm:"type:int unsigned not null;autoUpdateTime;comment:更新时间"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

const (
	// AccessTypeDefault 用户访问量
	AccessTypeDefault = 1
)
