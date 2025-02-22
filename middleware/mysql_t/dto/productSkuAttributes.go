package dto

import "gotest/common/models"

// ProductSkuAttributes SKU属性关联表
type ProductSkuAttributes struct {
	models.Model
	SkuID            uint `gorm:"type:int unsigned not null;index;comment:SKU ID" json:"skuId"`
	AttributeID      uint `gorm:"type:int unsigned not null;index;comment:属性ID" json:"attributeId"`
	AttributeValueID uint `gorm:"type:int unsigned not null;index;comment:属性值ID" json:"attributeValueId"`
}
