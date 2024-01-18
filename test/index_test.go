package main

import (
	"fmt"
	"github.com/google/uuid"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/log/zap_log"
	"gotest/common/utils"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// TestRandom 测试随机数
func TestRandom(t *testing.T) {
	array := []string{}
	for i := 0; i < 1000; i++ {
		array = append(array, utils.NewRandom().OrderSn())
	}

	testMap := map[string]string{}
	for i := 0; i < len(array); i++ {
		key := array[i]
		if testMap[key] != "" {
			zap_log.Logger.Debug("出现重复键:" + key + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		zap_log.Logger.Debug("key:" + key)
	}
	zap_log.Logger.Debug("完成")
}

// TestTime 测试时间
func TestTime(t *testing.T) {
	nowTime := time.Now()
	zap_log.Logger.Debug(strconv.Itoa(int(nowTime.Unix())))
	zap_log.Logger.Debug(strconv.Itoa(int(nowTime.UnixMilli())))
	zap_log.Logger.Debug(strconv.Itoa(int(nowTime.UnixNano())))
	zap_log.Logger.Debug(strconv.Itoa(int(nowTime.UnixMicro())))
}

// TestTime 测试时间
func TestArray(t *testing.T) {
	a := make([]int, 5)

	a[0] = 1
	a[1] = 2
	//a[2] = 3
	//a[3] = 4
	//a = append(a, 1)
	//a = append(a, 2)
	//a = append(a, 3)
	//a = append(a, 4)
	fmt.Println("a:", len(a))
}

// TestAutoTime 测试gorm 自动更新时间
func TestAutoTime(t *testing.T) {
	user := models.User{
		AdminId:  1,
		ParentId: 3,
	}
	database.DB.Create(&user)
	fmt.Println(user.UpdatedAt)
	fmt.Println(user.CreatedAt)
}

// TestUUID	 测试UUID
func TestUUID(t *testing.T) {
	newUUID := uuid.New().ID()
	zap_log.Logger.Debug(strconv.Itoa(int(newUUID)))
	uuidArray := []uint32{}
	for i := 0; i < 1000; i++ {
		uuidArray = append(uuidArray, uuid.New().ID())
	}
	testMap := map[uint32]string{}
	for i := 0; i < len(uuidArray); i++ {
		key := uuidArray[i]
		if testMap[key] != "" {
			zap_log.Logger.Debug("出现重复键:" + strconv.Itoa(int(key)) + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		zap_log.Logger.Debug("key:" + strconv.Itoa(int(key)))
	}
	zap_log.Logger.Debug("完成")
}

// TestRandomPrice 随机单价
func TestRandomPrice(t *testing.T) {
	currentPrice := 50.00
	mode := 2
	var resultPrice float64 = 0
	if mode == 1 {
		resultPrice = currentPrice * (1 + float64(randNum(1, 50))/10000)
	} else {
		resultPrice = currentPrice * float64(randNum(9950, 9999)) / 10000
	}
	fmt.Println("data:", resultPrice)
}

func randNum(m, n int) int {
	// 设置随机数种子
	rand.NewSource(time.Now().UnixNano())
	// 生成大于m且小于n的随机整数
	result := rand.Intn(n-m) + m
	return result
}
