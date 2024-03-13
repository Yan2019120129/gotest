package dto

import "gotest/common/models"

type AdminUserInfo struct {
	models.Model
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
