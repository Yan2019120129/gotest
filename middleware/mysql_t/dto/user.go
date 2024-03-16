package dto

import "gotest/common/models"

// UserInfo  用户信息
type UserInfo struct {
	models.Model
	AdminUserId int            `json:"adminUserId"`
	Telephone   string         `json:"telephone"`
	Sex         int            `json:"sex"`
	Birthday    int            `json:"birthday"`
	AdminUser   *AdminUserInfo `json:"adminUser" gorm:"embedded:foreignKey:AdminUserId;"`
	//AdminUser *models.AdminUser `json:"adminUser" gorm:"foreignKey:AdminUserId;"`
}
