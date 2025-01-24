package dto

import (
	"gotest/common/models"
)

const (
	ProductAttrsSkuStatusAvailable int8 = 10 // 上架
	ProductAttrsSkuStatusRemoved   int8 = -1 // 下架
)

// ProductSku 产品SKU
type ProductSku struct {
	models.Model
	ProductID     uint    `gorm:"type:int unsigned not null;index;comment:产品ID" json:"productId"`
	Name          string  `gorm:"type:varchar(512) not null;comment:SKU名称" json:"name"`
	Image         string  `gorm:"type:varchar(255) not null;comment:商品图片" json:"image"`
	Stock         uint    `gorm:"type:int unsigned not null;default:1000;comment:库存" json:"stock"`
	LockStock     uint    `gorm:"type:int unsigned not null;default:0;comment:锁定库存" json:"lockStock"`
	Sales         uint    `gorm:"type:int unsigned not null;default:0;comment:销量" json:"sales"`
	Money         float64 `gorm:"type:decimal(12, 2) not null;default:100;comment:原价" json:"money"`
	DiscountMoney float64 `gorm:"type:decimal(12, 2) not null;default:0;comment:折扣价" json:"discountMoney"`
	Status        int8    `gorm:"type:tinyint not null;default:10;index;comment:状态(10:上架,-1:下架)" json:"status"`
	Desc          string  `gorm:"type:text;comment:描述" json:"desc"`
}

type ProductAndSku struct {
	ProductSku
	ProductSkuAttributes []*ProductSkuAttributes `gorm:"foreignKey:SkuID" json:"productSkuAttributes"`
}

func (ProductAndSku) TableName() string {
	return "product_sku"
}
