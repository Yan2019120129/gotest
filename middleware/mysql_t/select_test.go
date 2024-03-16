package mysql_t

import (
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"gotest/middleware/mysql_t/dto"
	"testing"
)

// TestSelectBelongsTo 一对一查找方式
func TestSelectBelongsTo(t *testing.T) {
	// 查询用户和管理员
	userInfo := &models.User{}
	err := database.DB.Model(&models.User{}).Preload("AdminUser").Where("id = ?", 2).Find(userInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("userInfo", userInfo))
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("userInfo.AdminUser", userInfo.AdminUser))
}

// TestSelectBelongTo 测试一对一查询 并自动映射到自定义结构体测试
func TestSelectBelongTo(t *testing.T) {
	dtoUserInfo := &dto.UserInfo{}
	err := database.DB.Model(&models.User{}).
		//Preload("User").
		//Preload("User.AdminUser").
		Preload("AdminUser").
		Where("id = ?", 2).
		Find(dtoUserInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("dtoUserInfo", dtoUserInfo))
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("dtoUserInfo.AdminUser", dtoUserInfo.AdminUser))

}

// TestSelectHasOne 测试一对一关系
func TestSelectHasOne(t *testing.T) {
	// 查询管理员和用户
	userInfo := &dto.UserInfo{}
	err := database.DB.Model(&models.User{}).
		Preload("UserInfo").
		Preload("UserInfo.AdminUserInfo").
		Where("id = ?", 1).Take(userInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("adminUserInfo", userInfo))
	logs.Logger.Info(logs.LogMsgApp, zap.Reflect("adminUserInfo.Users", userInfo.AdminUser))
}
