package main

import (
	"gorm.io/gorm"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"strconv"
)

// ProductAttrsSkuList 产品属性SKU列表
type ProductAttrsSkuList struct {
	Ids  []string //	属性值ID列表
	Name []string //	属性值名称
}

// ProductAttrsKey 产品属性Key
type ProductAttrsKey struct {
	models.ProductAttrsKey
	Values []*models.ProductAttrsVal `gorm:"foreignKey:KeyId;"`
}

func GetProductAttrsList(tx *gorm.DB, productId uint) []*ProductAttrsKey {
	attrsList := make([]*ProductAttrsKey, 0)
	mode := database.DB
	if tx != nil {
		mode = tx
	}

	mode.Model(&models.ProductAttrsKey{}).Where("product_id = ?", productId).
		Preload("Values").
		Find(&attrsList)
	return attrsList
}

// ProductAttrsGenerateSkuList 产品属性生成Sku
func ProductAttrsGenerateSkuList(tx *gorm.DB, productId uint) []ProductAttrsSkuList {
	attrs := GetProductAttrsList(tx, productId)
	skuList := make([]ProductAttrsSkuList, 0)

	for i := 0; i < len(attrs[0].Values); i++ {
		sku := generateDescartesSkuList(
			attrs,
			ProductAttrsSkuList{
				Ids:  []string{strconv.FormatInt(int64(attrs[0].Values[i].ID), 10)},
				Name: []string{attrs[0].Values[i].Name},
			},
			1)
		skuList = append(skuList, sku...)
	}
	return skuList
}

// generateDescartesSkuList 产品属性笛卡尔积生成
func generateDescartesSkuList(attrsKeyList []*ProductAttrsKey, sep ProductAttrsSkuList, index int) []ProductAttrsSkuList {
	skuList := make([]ProductAttrsSkuList, 0)

	//	如果只有一位属性,  直接返回
	if len(attrsKeyList) < 2 {
		return []ProductAttrsSkuList{sep}
	}

	sepTmp := sep
	for i := 0; i < len(attrsKeyList[index].Values); i++ {
		sep = sepTmp
		sep.Ids = append(sep.Ids, strconv.FormatInt(int64(attrsKeyList[index].Values[i].ID), 10))
		sep.Name = append(sep.Name, attrsKeyList[index].Values[i].Name)

		if len(attrsKeyList)-1 == index {
			skuList = append(skuList, sep)
		} else {
			skuList = append(skuList, generateDescartesSkuList(attrsKeyList, sep, index+1)...)
		}
	}

	return skuList
}
