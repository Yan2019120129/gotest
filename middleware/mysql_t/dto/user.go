package dto

import "gotest/common/models"

// UserInfo  用户信息
type UserInfo struct {
	models.Model
	ParentId    int            `json:"parentId"`
	AdminUserId int            `json:"adminUserId"`
	Telephone   string         `json:"telephone"`
	Sex         int            `json:"sex"`
	Birthday    int            `json:"birthday"`
	AdminUser   *AdminUserInfo `json:"adminUser" gorm:"foreignKey:AdminUserId"`
}

func (u *UserInfo) TableName() string {
	return "user"
}
