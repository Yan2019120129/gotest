package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"time"
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

func main() {
	// 初始化配置
	currentTime := time.Now().Unix()
	userInfo := new(User)
	err := gofakeit.Struct(userInfo)
	userInfo.UpdatedAt = int(currentTime)
	userInfo.CreatedAt = int(currentTime)
	if err != nil {
		panic(err)
	}

	fmt.Println("userInfo:", userInfo)
}
