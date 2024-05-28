package main

import (
	"gorm.io/gorm"
	"gotest/common/models"
	"gotest/common/utils"
	"strings"
)

// AttrsAttr 产品属性信息
type AttrsAttr struct {
	models.ProductAttrsKey
	Value []*models.ProductAttrsVal `gorm:"foreignKey:KeyId"`
}

func (AttrsAttr) TableName() string {
	return "product_attrs_key"
}

// InsertProduct 插入爬取的产品
func InsertProduct(tx *gorm.DB, settingAdminId, categoryId uint, productAttr *ProductAttr) error {
	if err := tx.Transaction(func(tx *gorm.DB) error {
		discount := utils.Round((productAttr.GetOriginalPrice()-productAttr.GetCurrentPrice())/productAttr.GetOriginalPrice(), 2)
		productInfo := models.Product{
			AdminId:    settingAdminId,
			CategoryId: categoryId,
			Name:       productAttr.Title,
			Images:     productAttr.Images,
			Money:      productAttr.GetCurrentPrice(),
			Discount:   discount,
			Type:       2,
			Desc:       productAttr.Describe,
		}
		// 插入产品
		if err := tx.Create(&productInfo).Error; err != nil {
			return err
		}

		if len(productAttr.Style) > 0 {
			productKeyList := make([]*AttrsAttr, 0)
			for key, values := range productAttr.Style {
				attrsKeyInfo := &AttrsAttr{
					ProductAttrsKey: models.ProductAttrsKey{
						Model:     gorm.Model{},
						ProductId: productInfo.ID,
						Name:      key,
					},
				}
				for _, value := range values {
					attrsKeyInfo.Value = append(attrsKeyInfo.Value, &models.ProductAttrsVal{
						Name: value,
					})
				}
				productKeyList = append(productKeyList, attrsKeyInfo)
			}

			// 插入产品健值属性
			if err := tx.Create(&productKeyList).Error; err != nil {
				return err
			}

			generateSkuList := ProductAttrsGenerateSkuList(tx, productInfo.ID)

			productSkuList := make([]*models.ProductAttrsSku, 0)
			for _, skuInfo := range generateSkuList {
				productSkuList = append(productSkuList, &models.ProductAttrsSku{
					ProductId: productInfo.ID,
					Vals:      strings.Join(skuInfo.Ids, ","),
					Name:      strings.Join(skuInfo.Name, "."),
					Money:     productAttr.GetCurrentPrice(),
					Discount:  discount,
				})
			}

			// 插入产品sku属性
			if err := tx.Create(&productSkuList).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
