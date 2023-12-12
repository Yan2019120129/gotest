package test

import (
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gotest/my_frame/config/database/mysql"
	esearch "gotest/my_frame/config/elasticsearch"
	"gotest/my_frame/models"
	"log"
	"time"
)

// TestTime 测试gorm 自动更新时间
func TestTime() {
	product := models.WalletAssets{
		Name: "test",
	}
	mysql.Db.Create(&product)
	fmt.Println(product.UpdatedAt)
	fmt.Println(product.CreatedAt)
}

// TestMap 测试map条件查询
func TestMap() {
	var product []models.WalletAssets
	var field = map[string]interface{}{"name": "test"}
	mysql.Db.Where(field).Find(&product)
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
	err := mysql.Db.Transaction(func(tx *gorm.DB) error {
		for _, param := range params {
			var tempMap []map[string]interface{}
			mysql.Db.Table(param.Table).Where("admin_id = ?", 1).Find(&tempMap)

			for _, temp := range tempMap {
				// 遍历条件
				var field = map[string]interface{}{}
				for _, v := range param.WhereField {
					if val, ok := temp[v]; ok {
						field[v] = val
					}
				}

				var outcome map[string]interface{}
				mysql.Db.Table(param.Table).Where(field).Where("admin_id = ?", param.AdminId).Find(&outcome)
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
	err := mysql.Db.Transaction(func(tx *gorm.DB) error {
		for _, param := range params {

			// 获取超级管理员数据
			var tempMap []map[string]interface{}
			mysql.Db.Table(param.Table).Where("admin_id = ?", 1).Find(&tempMap)

			// 如果没有管理员数据则不进行下一步
			if len(tempMap) == 0 {
				continue
			}

			// 清空商户管理员数据
			mysql.Db.Table(param.Table).Delete("admin_id", param.AdminId)

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
	result := mysql.Db.AutoMigrate(&models.WalletAssets{})
	if result.Error() != "" {
		log.Panicln(result.Error())
	}
}

// InsertMysql 向表添加数据
func InsertMysql() {
	// 给我列举一个gorm 子查询例子，用请用中文回答我的问题
	var product models.Product

	mysql.Db.Where("id = ?", 1).Find(&product)

	fmt.Println(product)

}

// CreateIndex 创建缩影
func CreateIndex(index string) {
	do, err := esearch.ES.Clint.Indices.Create(index).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result", do.Index)
	fmt.Println("result：", do.Acknowledged)
	fmt.Println("result：", do.ShardsAcknowledged)
}

// DeleteIndex 删除索引
func DeleteIndex(index string) {
	do, err := esearch.ES.Clint.Indices.Delete(index).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do.Shards_)
	fmt.Println("result:", do.Acknowledged)
}

// FindIndexAll 获取节点信息
func FindIndexAll() {
	do, err := esearch.ES.Clint.Indices.GetTemplate().Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", do)
}

// InsertDocument 插入文档
func InsertDocument() {

}
