package dto

import (
	"gotest/common/models"
)

type AdminUserInfo struct {
	models.Model
	Username string       `json:"username"`
	Nickname string       `json:"nickname"`
	Email    string       `json:"email"`
	Avatar   string       `json:"avatar"`
	Users    *[]*UserInfo `json:"users" gorm:"foreignKey:AdminUserId;"`
}

func (a *AdminUserInfo) TableName() string {
	return "admin_user"
}
