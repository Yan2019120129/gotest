package mysql_t

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
	"gotest/middleware/mysql_t/dto"
	"os"
	"strings"
	"testing"
)

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

func TestInserterAttr(t *testing.T) {
	//attributes := []*dto.AttributesAndValue{
	//	{Attribute: dto.Attribute{Name: "color", Type: 1}, AttributeValues: []*dto.AttributeValue{
	//		{Name: "red"},
	//		{Name: "yellow"},
	//	}},
	//	{Attribute: dto.Attribute{Name: "size", Type: 1}, AttributeValues: []*dto.AttributeValue{
	//		{Name: "M"},
	//		{Name: "L"},
	//		{Name: "XL"},
	//	}},
	//}
	//
	//database.DB.Create(&attributes)
	attributes := make([]*dto.AttributesAndValue, 0)
	database.DB.Or("name = ?", "color").
		Or("name = ?", "size").
		Preload("AttributeValues").
		Find(&attributes)
	productAndSkus := make([]*dto.ProductAndSku, 0)
	productSku := dto.ProductSku{
		ProductID:     1,
		Name:          gofakeit.Name(),
		Image:         gofakeit.ImageURL(100, 200),
		Stock:         uint(gofakeit.Number(0, 10)),
		LockStock:     uint(gofakeit.Number(0, 10)),
		Sales:         uint(gofakeit.Number(0, 10)),
		Money:         50,
		DiscountMoney: 0,
		Status:        0,
		Desc:          "",
	}
	productAndSku := &dto.ProductAndSku{}
	for _, attribute := range attributes {
		for _, v := range attribute.AttributeValues {
			productAndSku.ProductSkuAttributes = append(productAndSku.ProductSkuAttributes, &dto.ProductSkuAttributes{AttributeID: attribute.Id, AttributeValueID: v.Id})
			productAndSkus = append(productAndSkus, productAndSku)
		}
	}
	database.DB.Create(&productAndSkus)
}
