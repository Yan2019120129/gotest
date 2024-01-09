package mysql_test

import (
	"go.uber.org/zap"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logger"
	"gotest/common/utils"
)

// TestWhereId 测试where 条件
func TestWhereId() {
	id := 24587
	userInfo := &models.User{}
	if result := database.DB.Where(id).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereIds 测试where 条件
func TestWhereIds() {
	ids := []int{24587, 24588}
	userInfo := &models.User{}
	if result := database.DB.Where(ids).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}

// TestWhereOtherIds 测试where 条件，测试结果不可行，会自动映射为主键ID
func TestWhereOtherIds() {
	adminIds := []int{1, 2}
	userInfo := &models.User{}
	if result := database.DB.Where(adminIds).Find(userInfo); result.Error != nil {
		logger.Logger.Warn("错误信息", zap.Error(result.Error))
		return
	}
	logger.Logger.Info("信息", zap.String("data", utils.ObjToString(userInfo)))
}
