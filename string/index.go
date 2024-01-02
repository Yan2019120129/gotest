package main

import (
	"fmt"
	"strings"
)

func main() {
	findValue()
}

// findValue 查找字符字符串对应的值
func findValue() {
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
