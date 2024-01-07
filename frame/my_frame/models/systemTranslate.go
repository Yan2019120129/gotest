package models

// Translate 系统语言翻译表
type Translate struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	Lang      string `gorm:"type:varchar(60) not null;index:idx_lang;comment:语言标识"`
	Name      string `gorm:"type:varchar(60) not null;comment:名称"`
	Type      int    `gorm:"type:tinyint unsigned not null default 1;default:1;comment:类型 1系统翻译 2前台翻译"`
	Field     string `gorm:"type:varchar(60) not null;uniqueIndex;comment:建铭"`
	Value     string `gorm:"type:text;comment:键值"`
	CreatedAt int    `gorm:"type:int unsigned not null;comment:创建时间"`
}

const (
	// TranslateTypeDefault 系统翻译
	TranslateTypeDefault = 1

	// TranslateTypeHome 前台翻译
	TranslateTypeHome = 2
)
