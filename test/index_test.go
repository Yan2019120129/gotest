package main

import (
	"fmt"
	"github.com/google/uuid"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logger"
	"gotest/common/utils"
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
			logger.Logger.Debug("出现重复键:" + key + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		logger.Logger.Debug("key:" + key)
	}
	logger.Logger.Debug("完成")
}

// TestTime 测试时间
func TestTime(t *testing.T) {
	nowTime := time.Now()
	logger.Logger.Debug(strconv.Itoa(int(nowTime.Unix())))
	logger.Logger.Debug(strconv.Itoa(int(nowTime.UnixMilli())))
	logger.Logger.Debug(strconv.Itoa(int(nowTime.UnixNano())))
	logger.Logger.Debug(strconv.Itoa(int(nowTime.UnixMicro())))
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
	logger.Logger.Debug(strconv.Itoa(int(newUUID)))
	uuidArray := []uint32{}
	for i := 0; i < 1000; i++ {
		uuidArray = append(uuidArray, uuid.New().ID())
	}
	testMap := map[uint32]string{}
	for i := 0; i < len(uuidArray); i++ {
		key := uuidArray[i]
		if testMap[key] != "" {
			logger.Logger.Debug("出现重复键:" + strconv.Itoa(int(key)) + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		logger.Logger.Debug("key:" + strconv.Itoa(int(key)))
	}
	logger.Logger.Debug("完成")
}
