package string_t

import (
	"fmt"
	"strconv"
	"strings"
)

// FindValue 查找字符字符串对应的值
func FindValue() {
	jsonStr := `{
		"op": "subscribe",
		"args": [{
			"channel": "option-trades",
			"instType": "OPTION",
			"instFamily": "BTC-USD"
		}]
	}`

	index := strings.Index(jsonStr, "channel")
	fmt.Println("channel:", index)
	fmt.Println("data:", jsonStr[index+1:])
}

// TestStrconv  测试strconv.FormatInt 方法,将字符串转换为int 十进制类型
func TestStrconv() {
	formatIntValue := strconv.FormatInt(15, 10)
	fmt.Printf("类型：%T \n", formatIntValue)
	fmt.Println("值：", formatIntValue)
}

// MultipleDataString 测试传入多数据的情况
func MultipleDataString(s ...string) {
	fmt.Printf("类型:%T", s)
	fmt.Println("值:", s)
	for i, v := range s {
		fmt.Printf("%v值:%v\n", i, v)
	}
}

// MultipleDataInt 测试传入多数据的情况
func MultipleDataInt(data ...int) {
	fmt.Printf("类型:%T", data)
	fmt.Println("值:", data)
	for i, v := range data {
		fmt.Printf("%v值:%v\n", i, v)
	}
}
