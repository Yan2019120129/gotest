package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"gotest/my_frame/models"
	"gotest/my_frame/module/gorm/database"
	"gotest/test/test_init"
	"time"
)

var Success string

func main() {
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

func init() {
	Success = test_init.Ok
}
