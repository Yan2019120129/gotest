package models

import (
	"gorm.io/gorm"
)

const (
	ProductAttrsSkuStatusActivate = 10
)

// ProductAttrsSku 产品属性SKU
type ProductAttrsSku struct {
	gorm.Model
	ParentId  uint    `gorm:"type:int unsigned not null;comment:父级ID"`
	ProductId uint    `gorm:"type:int unsigned not null;default:0;comment:产品ID"`
	Vals      string  `gorm:"type:varchar(255) not null;default:'';comment:属性值ID,用逗号分隔"`
	Name      string  `gorm:"type:varchar(512) not null;default:'';comment:SKU名称"`
	Image     string  `gorm:"type:varchar(255) not null;default:'';comment:商品图片"`
	Stock     uint    `gorm:"type:int unsigned not null;default:1000;comment:库存量"`
	Sales     uint    `gorm:"type:int unsigned not null;default:0;comment:销售量"`
	Money     float64 `gorm:"type:decimal(12, 2) not null;default:100;comment:标价"`
	Discount  float64 `gorm:"type:decimal(8,4) not null;default:0;comment:折扣"`
	Data      string  `gorm:"type:varchar(255) not null;default:'';comment:数据"`
	Status    int     `gorm:"type:tinyint not null;default:10;comment:状态 -1下架｜10上架"`
}
