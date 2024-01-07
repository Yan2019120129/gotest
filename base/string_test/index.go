package string_test

import (
	"fmt"
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
