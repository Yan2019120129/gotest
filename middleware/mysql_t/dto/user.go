package dto

import (
	"github.com/brianvoe/gofakeit/v6"
	"gotest/common/models"
)

// UserInfo  用户信息
type UserInfo struct {
	models.Model
	ChannelId   uint    `gorm:"primaryKey;type:int unsigned not null;comment:渠道ID" json:"channelId"`
	AdminId     int     `json:"adminUserId"`
	ParentId    int     `json:"parentId"`
	Telephone   string  `json:"telephone"`
	Sex         int     `json:"sex"`
	Birthday    int     `json:"birthday"`
	Username    string  `json:"username"`
	Nickname    string  `json:"nickname"`
	Email       string  `json:"email"`
	Avatar      string  `json:"avatar"`
	Password    string  `json:"password"`
	SecurityKey string  `json:"securityKey"`
	Money       float64 `json:"money"`
}

func (u *UserInfo) TableName() string {
	return "user"
}

// GetUserInfoDefault 获取默认信息
func GetUserInfoDefault() *UserInfo {
	return &UserInfo{
		AdminId:     0,
		ParentId:    0,
		Username:    gofakeit.Name(),
		Nickname:    gofakeit.Name(),
		Email:       gofakeit.Email(),
		Telephone:   gofakeit.Phone(),
		Avatar:      gofakeit.ImageURL(200, 100),
		Sex:         gofakeit.RandomInt([]int{1, 2}),
		Password:    gofakeit.Password(true, true, true, false, false, 10),
		SecurityKey: gofakeit.Password(true, true, true, false, false, 10),
		Money:       gofakeit.Float64Range(100, 100000),
	}
}

//// SetUserInfoParent 设置关联信息
//func (u *UserInfo) SetUserInfoParent(userInfo *UserInfo) *UserInfo {
//	u.Users = append(u.Users, userInfo)
//	return u
//}
//
//// SetAdminUserInfo 设置管理员信息
//func (u *UserInfo) SetAdminUserInfo(adminUserInfo *AdminUserInfo) *UserInfo {
//	u.AdminUser = adminUserInfo
//	return u
//}
