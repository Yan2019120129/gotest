package indexserver

import (
	"my-frame/models"
	"my-frame/module/gorm/database"
)

// Login 登录接口
func Login() (interface{}, error) {
	userInfo := &models.AdminUser{}
	result := database.DB.First(userInfo)
	return userInfo, result.Error
}
