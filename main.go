package main

import (
	"fmt"
	"strings"
)

func main() {
	// 执行计算并保留两位小数
	bwSumTmp := "4.2\n"
	bwSumTmp = strings.ReplaceAll(bwSumTmp, "\n", "")
	fmt.Println(bwSumTmp)
}
