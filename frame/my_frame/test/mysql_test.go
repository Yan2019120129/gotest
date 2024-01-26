package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"gotest/frame/my_frame/models"
	"gotest/frame/my_frame/module/gorm/database"
	"log"
	"testing"
	"time"
)

// TestConnectTable 测试连表同步
func TestConnectTable(t *testing.T) {
	var product []models.WalletAssets
	var field []map[string]interface{}
	database.DB.Where(field).Find(&product)
	fmt.Println(product)
}

// TestLogin 测试登录接口
func TestLogin(t *testing.T) {

}

// TestTime 测试gorm 自动更新时间
func TestTime(t *testing.T) {
	user := models.User{
		AdminId:  1,
		ParentId: 3,
	}
	database.DB.Create(&user)
	fmt.Println(user.UpdatedAt)
	fmt.Println(user.CreatedAt)
}

// TestMap 测试map条件查询
func TestMap(t *testing.T) {
	var product []models.WalletAssets
	var field = map[string]interface{}{"name": "test"}
	database.DB.Where(field).Find(&product)
	fmt.Println(product)
}

type SynchronousParams struct {
	Table      string   `json:"table"`      // 表名
	AdminId    int      `json:"admin_id"`   // 复制的表名
	WhereField []string `json:"whereField"` // 条件字段
}

type CopyParams struct {
	Table   string `json:"table"`    // 表名
	AdminId int    `json:"admin_id"` // 复制的表名
}

// Synchronous 同步
func Synchronous(params []*SynchronousParams) error {

	nowTime := int(time.Now().Unix())
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		for _, param := range params {
			var tempMap []map[string]interface{}
			database.DB.Table(param.Table).Where("admin_id = ?", 1).Find(&tempMap)

			for _, temp := range tempMap {
				// 遍历条件
				var field = map[string]interface{}{}
				for _, v := range param.WhereField {
					if val, ok := temp[v]; ok {
						field[v] = val
					}
				}

				var outcome map[string]interface{}
				database.DB.Table(param.Table).Where(field).Where("admin_id = ?", param.AdminId).Find(&outcome)
				if outcome != nil {
					continue
				}

				temp["id"], temp["admin_id"], temp["created_at"] = 0, param.AdminId, nowTime
				if temp["updated_at"] != nil {
					temp["updated_at"] = nowTime
				}

				result := tx.Table(param.Table).Create(&temp)
				if result.Error != nil {
					log.Println(result.Error)
					return result.Error
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Copy 复制
func Copy(params []*CopyParams) error {
	nowTime := int(time.Now().Unix())
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		for _, param := range params {

			// 获取超级管理员数据
			var tempMap []map[string]interface{}
			database.DB.Table(param.Table).Where("admin_id = ?", 1).Find(&tempMap)

			// 如果没有管理员数据则不进行下一步
			if len(tempMap) == 0 {
				continue
			}

			// 清空商户管理员数据
			database.DB.Table(param.Table).Delete("admin_id", param.AdminId)

			// 更新查询数据准备插入
			for _, temp := range tempMap {
				temp["id"], temp["admin_id"], temp["created_at"] = 0, param.AdminId, nowTime
				if temp["updated_at"] != nil {
					temp["updated_at"] = nowTime
				}
			}

			// 插入数据
			result := tx.Table(param.Table).Create(tempMap)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateTable 创建表
func CreateTable() {
	result := database.DB.AutoMigrate(&models.WalletAssets{})
	if result.Error() != "" {
		log.Panicln(result.Error())
	}
}

// InsertMysql 向表添加数据
func TestInsertMysql(t *testing.T) {
	// 给我列举一个gorm 子查询例子，用请用中文回答我的问题
	user := &models.User{
		AdminId:     gofakeit.Number(1, 1),
		ParentId:    gofakeit.Number(1, 1),
		UserName:    gofakeit.Name(),
		NickName:    gofakeit.Name(),
		Email:       gofakeit.Email(),
		Telephone:   gofakeit.Phone(),
		Avatar:      gofakeit.ImageURL(100, 300),
		Sex:         gofakeit.Number(1, 2),
		Birthday:    int(time.Now().Unix()),
		Password:    gofakeit.Password(true, true, true, true, true, 10),
		SecurityKey: gofakeit.Password(true, true, true, true, true, 10),
		Money:       8000,
		Type:        1,
		Status:      10,
		Data:        gofakeit.Phrase(),
		Desc:        gofakeit.Phrase(),
		UpdatedAt:   int(time.Now().Unix()),
		CreatedAt:   int(time.Now().Unix()),
	}
	if result := database.DB.Create(user); result.Error != nil {
		panic(result.Error)
		return
	}
}
