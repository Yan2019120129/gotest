package dto

import "gotest/common/models"

// AttributeValue 属性值
type AttributeValue struct {
	models.Model
	AttributeID uint   `gorm:"type:int unsigned not null;index;comment:属性ID" json:"attributeId"`
	Name        string `gorm:"type:varchar(255) not null;comment:名称" json:"name"`
}
