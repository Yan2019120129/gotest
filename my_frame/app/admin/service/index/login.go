package indexserver

import (
	"gotest/my_frame/models"
	"gotest/my_frame/module/gorm/database"
)

// Login 登录接口
func Login() (interface{}, error) {
	userInfo := &models.AdminUser{}
	result := database.DB.First(userInfo)
	return userInfo, result.Error
}
