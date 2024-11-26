package mysql_t

import (
	"fmt"
	"go.uber.org/zap"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"gotest/middleware/mysql_t/dto"
	"os"
	"strings"
	"testing"
)

// CreateTable 创建表
func TestCreateTable(t *testing.T) {
	err := database.DB.AutoMigrate(&dto.UserInfo{}, &dto.AdminUserInfo{})
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}

}

// GetColumn 获取字段信息
func TestGetColumn(t *testing.T) {
	types, err := database.DB.Migrator().ColumnTypes(&dto.UserInfo{})
	if err != nil {
		return
	}

	for _, v := range types {
		logs.Logger.Info(logs.LogMsgTest, zap.String("column", v.Name()))
	}
}

// CreateTable 插入数据
func TestInserter(t *testing.T) {
	userList := make([]*dto.UserInfo, 0)
	for i := 0; i < 10; i++ {
		userList = append(userList, dto.GetUserInfoDefault())
	}
	adminUserInfo := dto.GetAdminInfoDefault().SetUser(userList...)
	err := database.DB.Create(adminUserInfo).Error
	if err != nil {
		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
		return
	}
}

//
//// CreateTable 插入数据,使用自定义结构体
//func TestInserterCustomize(t *testing.T) {
//	userInfo := dto.GetUserInfoDefault().SetUserInfoParent(dto.GetUserInfoDefault()).SetUserInfoParent(dto.GetUserInfoDefault()).SetAdminUserInfo(dto.GetAdminInfoDefault())
//	err := database.DB.Create(userInfo).Error
//	if err != nil {
//		logs.Logger.Error(logs.LogMsgApp, zap.Error(err))
//		return
//	}
//}

var sqlPaths = []string{
	"./common/file/productsql/product.sql",
	"./common/file/productsql/product_attrs_key.sql",
	"./common/file/productsql/product_attrs_sku.sql",
	"./common/file/productsql/product_attrs_val.sql",
}

// TestInserterBySql 通过sql 插入数据
func TestInserterBySql(t *testing.T) {
	sqlPath := "/home/programmer-yan/Documents/GoFile/gotest/common/file/productsql"
	dir, err := os.ReadDir(sqlPath)
	if err != nil {
		fmt.Println("err-------:", err)
		return
	}
	for _, d := range dir {
		if !d.IsDir() {
			path := sqlPath + "/" + d.Name()
			if strings.Contains(d.Name(), ".sql") {
				sqls, err := os.ReadFile(path)
				if err != nil {
					panic(err)
				}
				sqlArr := strings.Split(string(sqls), ";\n")
				for _, s := range sqlArr {
					if s != "" {
						if err = database.DB.Exec(s).Error; err != nil {
							fmt.Println("sql", s, err)
							return
						}
					}
				}
			}
		}
	}
}

// 导出sql 数据
func TestExportSql(t *testing.T) {
}
