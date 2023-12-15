package index

import (
	"gotest/my_frame/config/gorm/database"
	"gotest/my_frame/models"
)

// Login 登录接口
func Login() (interface{}, error) {
	userInfo := &models.AdminUser{}
	result := database.DB.First(userInfo)
	return userInfo, result.Error
}
