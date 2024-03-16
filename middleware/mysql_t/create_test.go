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
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}

// CreateTable 插入数据
func TestInserter(t *testing.T) {
	err := database.DB.Create(models.GetDefaultUser().SetAdminUser(models.GetDefaultAdminUser())).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}

// TestInsertAssociation 插入数据并关联Id
func TestAssociation(t *testing.T) {
	err := database.DB.Create(models.GetDefaultUser().SetAdminId(1)).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}
