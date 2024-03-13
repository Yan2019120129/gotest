package mysql_t

import (
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"testing"
)

// TestSelectBelongsTo 一对一查找方式
func TestSelectBelongsTo(t *testing.T) {
	// 查询用户和管理员
	//userInfo := &models.User{}
	userInfo := make(map[string]interface{})
	err := database.DB.Model(&models.User{}).Preload("AdminUser").Where("id = ?", 2).Find(userInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("userInfo", userInfo))
	//logs.Logger.Info(logs.LogMsgApp, zap.Reflect("userInfo.AdminUser", userInfo.AdminUser))

	//dtoUserInfo := &dto.UserInfo{}
	//err := database.DB.Model(&models.User{}).Where("id = ?", 2).Preload("AdminUser").Find(dtoUserInfo).Error
	//if err != nil {
	//	logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	//	return
	//}
	//logs.Logger.Info(logs.LogMsgApp, zap.Reflect("dtoUserInfo", dtoUserInfo))
	//logs.Logger.Info(logs.LogMsgApp, zap.Reflect("dtoUserInfo.AdminUser", dtoUserInfo.AdminUser))

	// 查询管理员和用户
	//adminUserInfo := &models.AdminUser{}
	//err := database.DB.Preload("User").Where("id = ?", 1).Find(adminUserInfo).Error
	//if err != nil {
	//	logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	//	return
	//}
	//logs.Logger.Info(logs.LogMsgApp, zap.Reflect("adminUserInfo", adminUserInfo))
	//logs.Logger.Info(logs.LogMsgApp, zap.Reflect("adminUserInfo.Users", adminUserInfo.User))

}
