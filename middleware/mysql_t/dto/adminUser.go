package dto

import (
	"github.com/brianvoe/gofakeit/v6"
	"gotest/common/models"
)

type AdminUserInfo struct {
	models.Model
	Users       []*UserInfo `json:"users" gorm:"foreignKey:AdminId;"`
	ParentId    int         `json:"parentId"`
	Username    string      `json:"username"`
	Nickname    string      `json:"nickName"`
	Email       string      `json:"email"`
	Avatar      string      `json:"avatar"`
	Password    string      `json:"password"`
	SecurityKey string      `json:"securityKey"`
	Money       float64     `json:"money"`
	Status      int         `json:"status"`
	Domains     string      `json:"domains"`
	ExpiredAt   int         `json:"ExpiredAt"`
}

func (a *AdminUserInfo) TableName() string {
	return "admin_user"
}

// GetAdminInfoDefault 获取自定义AdminUser 信息
func GetAdminInfoDefault() *AdminUserInfo {
	return &AdminUserInfo{
		Username:    gofakeit.Name(),
		Nickname:    gofakeit.Name(),
		Email:       gofakeit.Email(),
		Avatar:      gofakeit.ImageURL(200, 100),
		Password:    gofakeit.Password(true, true, true, false, false, 10),
		SecurityKey: gofakeit.Password(true, true, true, false, false, 10),
		Money:       gofakeit.Float64Range(100, 100000),
		Status:      gofakeit.RandomInt([]int{-2, -1, 10}),
		Domains:     gofakeit.Letter(),
	}
}

// SetUser  设置用户数据
func (a *AdminUserInfo) SetUser(users ...*UserInfo) *AdminUserInfo {
	a.Users = append(a.Users, users...)
	return a
}

// SetParentId  设置上级Id
func (a *AdminUserInfo) SetParentId(parentId int) *AdminUserInfo {
	a.ParentId = parentId
	return a
}
