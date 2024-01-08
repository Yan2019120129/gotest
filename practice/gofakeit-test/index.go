package gofakeit_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gotest/common/models"
)

// User 用户表
type User struct {
	Id          int     `fake:"-"`
	AdminId     int     `fake:"{randomint:[1]}"`
	ParentId    int     `fake:"{randomint:[1]}"`
	UserName    string  `fake:"{name}"`
	NickName    string  `fake:"{name}"`
	Email       string  `fake:"{email}"`
	Telephone   string  `fake:"{phone}"`
	Avatar      string  `fake:"{imageurl}"`
	Sex         int     `fake:"{randomint:[1,2]}"`
	Birthday    int     `fake:"{number:946656000,1104508800}"`
	Password    string  `fake:"{randomstring:['1234567890','qwertyuiopasdfghjklzxcvbnm','QWERTYUIOPASDFGHJKLZXCVBNM']}"`
	SecurityKey string  `fake:"{randomstring:['1234567890','qwertyuiopasdfghjklzxcvbnm','QWERTYUIOPASDFGHJKLZXCVBNM']}"`
	Money       float64 `fake:"{float64range:100,1000}"`
	Type        int     `fake:"{randomint:[-2,-1,10]}"`
	Status      int     `fake:"{randomint:[-2,-1,10]}"`
	Data        string  `fake:"{loremipsumsentence}"`
	UpdatedAt   int     `fake:"{day}"`
	CreatedAt   int     `fake:"{day}"`
}

// GenerateData 生成数据
func GenerateData() {
	account := models.Account{
		AccountNumber: gofakeit.Number(1, 100000),
		Address:       gofakeit.Address(),
		Age:           gofakeit.Number(18, 25),
		Balance:       gofakeit.Number(500, 2000),
		City:          gofakeit.City(),
		Email:         gofakeit.Email(),
		Employer:      gofakeit.Name(),
		Firstname:     gofakeit.FirstName(),
		Gender:        gofakeit.Gender(),
		Lastname:      gofakeit.LastName(),
		State:         gofakeit.State(),
	}
	fmt.Println("account：", account)
}
