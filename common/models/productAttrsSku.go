package models

import (
	"github.com/gomodule/redigo/redis"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"gorm.io/gorm"
	"strconv"
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

// GetStoreSkuId 获取店铺SkuId
func (_Sku *ProductAttrsSku) GetStoreSkuId(rdsConn redis.Conn, adminIds []uint) []*views.InputOptions {
	data := make([]*views.InputOptions, 0)

	// 获取管理员对应的店铺ID和父级ID都不为0的产品
	productList := make([]*Product, 0)
	database.Db.Model(&Product{}).
		Where("admin_id IN ?", adminIds).
		Where("store_id <> ?", 0).
		Where("parent_id <> ?", 0).
		Find(&productList)

	// 获取产品对应的Sku名称
	skuList := make([]*ProductAttrsSku, 0)
	for _, product := range productList {
		database.Db.
			Where("product_id = ?", product.ID).
			Find(&skuList)

		for _, sku := range skuList {
			data = append(data, &views.InputOptions{
				Label: "ID: { " + strconv.Itoa(int(product.ID)) + " }，productName: { " + product.Name + "}，skuName: { " + sku.Name + " } ",
				Value: sku.ID,
			})
		}
	}

	return data
}
