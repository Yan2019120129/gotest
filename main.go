package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gotest/models"
	"time"
)

// Product 结构体表示产品信息
type Product struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	dsn := "root:Aa123098..@tcp(127.0.0.1:3306)/basic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	nowTime := time.Now()
	user := models.User{
		AdminId:     1,
		ParentId:    0,
		UserName:    "wang",
		NickName:    "wangyang",
		Email:       "15564036382",
		Telephone:   "188494",
		Avatar:      "/public/images/logo.png",
		Sex:         1,
		Birthday:    1700998487,
		Password:    "13135",
		SecurityKey: "unlock",
		Money:       1000000,
		Type:        models.UserTypeDefault,
		Status:      models.UserStatusActive,
		Data:        "粗布麻衣生涯，包含诗书气自华",
		Desc:        "",
		UpdatedAt:   int(nowTime.Unix()),
		CreatedAt:   int(nowTime.Unix()),
	}

	// 手动事务
	//tx := db.Begin()
	//defer tx.Rollback()
	//tx.Create(&user)
	//panic(errors.New("打断更新"))
	//tx.Model(&models.User{}).Where("id=?", user.Id).Updates(models.User{
	//	UserName:  "li",
	//	NickName:  "lifu",
	//	Email:     "214656551",
	//	Telephone: "98765242",
	//})
	//
	//tx.Commit()

	// 自动事务
	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		if true {
			// 返回 nil 提交事务
			return gorm.ErrRecordNotFound
		}
		tx.Model(&models.User{}).Where("id=?", user.Id).Updates(models.User{
			UserName:  "li",
			NickName:  "lifu",
			Email:     "214656551",
			Telephone: "98765242",
		})
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		panic(err)
	}
}
