package models

import "gorm.io/gorm"

const (
	ProductAttrsValStatusActivate = 10
)

// ProductAttrsVal 产品属性值
type ProductAttrsVal struct {
	gorm.Model
	KeyId  uint   `gorm:"type:int unsigned not null;default:0;comment:属性名称ID"`
	Name   string `gorm:"type:varchar(191) not null;default:'';comment:属性值名称"`
	Data   string `gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status int    `gorm:"type:tinyint not null;default:10;comment:状态 -1禁用｜10启用"`
}
