package mysql_t

import (
	"go.uber.org/zap"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"gotest/middleware/mysql_t/dto"
	"testing"
)

// CreateTable 创建表
func TestCreateTable(t *testing.T) {
	err := database.DB.AutoMigrate(&dto.Attribute{}, &dto.AttributeValue{}, &dto.ProductSku{}, &dto.ProductSkuAttributes{})
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}

}
