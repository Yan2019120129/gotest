package mysql_t

import (
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"testing"
)

// CreateTable 创建表
func TestCreateTable(t *testing.T) {
	err := database.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductCategory{}, &models.WalletAssets{})
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}

// CreateTable 插入数据
func TestInserter(t *testing.T) {
	userList := make([]*models.User, 0)
	for i := 0; i < 10; i++ {
		userList = append(userList, models.GetDefaultUser())
	}
	adminUserInfo := models.GetDefaultAdminUser().SetUser(userList...)
	err := database.DB.Create(adminUserInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}

// TestInsertAssociation 插入数据并关联Id
func TestAssociation(t *testing.T) {
	//err := database.DB.Create(models.GetDefaultUser().SetAdminId(1)).Error
	//if err != nil {
	//	logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
	//	return
	//}

	err := database.DB.Create(models.GetDefaultUser().SetAdminId(1).SetParentId(2)).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}
