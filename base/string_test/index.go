package string_test

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
