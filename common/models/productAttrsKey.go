package models

import "gorm.io/gorm"

const (
	ProductAttrsKeyStatusActivate = 10
)

// ProductAttrsKey 产品属性名称
type ProductAttrsKey struct {
	gorm.Model
	ProductId uint   `gorm:"type:int unsigned not null;default:0;comment:产品ID"`
	Name      string `gorm:"type:varchar(191) not null;default:'';comment:属性名称"`
	Type      int    `gorm:"type:tinyint not null;default:1;comment:类型 1商品属性"`
	Data      string `gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status    int    `gorm:"type:tinyint not null;default:10;comment:状态 -1禁用｜10启用"`
}
