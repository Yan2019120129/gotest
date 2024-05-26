package models

import (
	"gorm.io/gorm"
)

const (
	ProductStatusActive  = 10
	ProductStatusDisable = -1

	// ProductTypeDefault 店铺商品类型
	ProductTypeDefault = 1
	// ProductTypeWholesale  批发商品类型
	ProductTypeWholesale = 2
)

// Product 产品表
type Product struct {
	gorm.Model
	AdminId    uint        `gorm:"type:int unsigned not null;comment:管理员ID"`
	ParentId   uint        `gorm:"type:int unsigned not null;comment:父级ID"`
	StoreId    uint        `gorm:"type:int unsigned not null;comment:店铺ID"`
	CategoryId uint        `gorm:"type:int unsigned not null;comment:类目ID"`
	AssetsId   uint        `gorm:"type:int unsigned not null;comment:资产ID"`
	Name       string      `gorm:"type:varchar(512) not null;comment:标题"`
	Images     GormStrings `gorm:"type:varchar(4096) not null;comment:商品图片"`
	Money      float64     `gorm:"type:decimal(20,6) not null;comment:标价"`
	Discount   float64     `gorm:"type:decimal(8,4) not null;comment:折扣"`
	Type       int         `gorm:"type:tinyint not null;default:1;comment:类型1店铺商品 2批发商品"`
	Sort       int         `gorm:"type:tinyint not null;default:99;comment:排序"`
	Sales      int         `gorm:"type:int unsigned not null;default:0;comment:销售量"`
	Status     int         `gorm:"type:tinyint not null;default:10;comment:状态-1禁用 10启用"`
	Data       string      `gorm:"type:text;comment:数据"`
	Desc       string      `gorm:"type:text;comment:描述"`
}
