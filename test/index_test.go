package test

import (
	"fmt"
	"github.com/google/uuid"
	"gotest/common/models"
	"gotest/common/module/gorm/database"
	"gotest/common/module/logs"
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
			logs.Logger.Debug("出现重复键:" + key + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		logs.Logger.Debug("key:" + key)
	}
	logs.Logger.Debug("完成")
}

// TestTime 测试时间
func TestTime(t *testing.T) {
	nowTime := time.Now()
	logs.Logger.Debug(strconv.Itoa(int(nowTime.Unix())))
	logs.Logger.Debug(strconv.Itoa(int(nowTime.UnixMilli())))
	logs.Logger.Debug(strconv.Itoa(int(nowTime.UnixNano())))
	logs.Logger.Debug(strconv.Itoa(int(nowTime.UnixMicro())))
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
		AdminUserId: 1,
		ParentId:    3,
	}
	database.DB.Create(&user)
	fmt.Println(user.UpdatedAt)
	fmt.Println(user.CreatedAt)
}

// TestUUID	 测试UUID
func TestUUID(t *testing.T) {
	newUUID := uuid.New().ID()
	logs.Logger.Debug(strconv.Itoa(int(newUUID)))
	uuidArray := []uint32{}
	for i := 0; i < 1000; i++ {
		uuidArray = append(uuidArray, uuid.New().ID())
	}
	testMap := map[uint32]string{}
	for i := 0; i < len(uuidArray); i++ {
		key := uuidArray[i]
		if testMap[key] != "" {
			logs.Logger.Debug("出现重复键:" + strconv.Itoa(int(key)) + "-" + strconv.Itoa(i))
			return
		}
		testMap[key] = strconv.Itoa(i)
		logs.Logger.Debug("key:" + strconv.Itoa(int(key)))
	}
	logs.Logger.Debug("完成")
}

// TestRandomPrice 随机单价
func TestRandomPrice(t *testing.T) {
	currentPrice := 100000.0
	mode := 2
	accuracy := 3
	randomPrice := 5.0
	denominator := 1.0
	for i := 0; i < accuracy; i++ {
		denominator = denominator * 10
	}
	var resultPrice float64 = 0
	switch mode {
	case 1:
		resultPrice = currentPrice + float64(randNum(1, 5))/denominator
	case 2:
		resultPrice = currentPrice - float64(randNum(1, 5))/denominator
	}
	fmt.Println("denominator", randomPrice*1/denominator)
	fmt.Println("resultPrice", resultPrice)
}

func randNum(m, n int) int {
	// 设置随机数种子
	rand.NewSource(time.Now().UnixNano())
	// 生成大于m且小于n的随机整数
	result := rand.Intn(n-m) + m
	return result
}

type People struct {
	Name string
	Age  int
}

// TestAddr 测试地址传递
func TestAddr(t *testing.T) {
	people := People{
		Name: "yan",
		Age:  18,
	}
	tempPeople := people
	tempPeopleOne := People{
		Name: "yan",
		Age:  18,
	}
	intValue := 1
	IntAddr(&intValue)
	fmt.Printf("int:%p\n", &intValue)
	PeopleAddr(people)
	fmt.Printf("people:%p\n", &people)
	fmt.Printf("peopleEqtempPeople:%v,p1:%p,p1:%p,p1Eqp2:%v\n", people == tempPeople, &people, &tempPeople, &people == &tempPeople)
	fmt.Printf("peopleEqtempPeopleOne:%v,p1:%p,p1:%p\n", people == tempPeopleOne, &people, &tempPeopleOne)

}

// IntAddr addr 测试int 类型地址传输区别
func IntAddr(i *int) {
	fmt.Printf("i:%p\n", &i)
}

// PeopleAddr addr 测试 PeopleAddr 类型地址传输区别
func PeopleAddr(p People) {
	fmt.Printf("PeopleAddr:%p\n", &p)
	fmt.Printf("PeopleAddr&*:%p\n", &p)
}
