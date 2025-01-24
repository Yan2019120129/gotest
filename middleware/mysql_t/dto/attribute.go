package dto

import (
	"gotest/common/models"
)

// Attribute 属性
type Attribute struct {
	models.Model
	Name string `gorm:"type:varchar(255) not null;comment:名称" json:"name"`
	Type int8   `gorm:"type:tinyint not null;default:1;comment:类型(1:默认)" json:"type"`
}

type AttributesAndValue struct {
	Attribute
	AttributeValues []*AttributeValue `gorm:"foreignKey:AttributeID" json:"attribute_values"`
}

func (Attribute) TableName() string {
	return "attribute"
}
