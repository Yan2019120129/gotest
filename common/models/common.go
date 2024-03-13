package models

type Model struct {
	Id        int `gorm:"type:int unsigned primary key;comment:主键;"`
	UpdatedAt int `gorm:"type:int unsigned not null;comment:更新时间"`
	CreatedAt int `gorm:"type:int unsigned not null;comment:创建时间"`
}
