// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"fmt"
	"gotest/common/utils"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(OrderRandomPrice(2367.33, 1, 2))
	}
}

// OrderRandomPrice 随机订单价格 1涨 2跌
func OrderRandomPrice(currentPrice float64, mode, accuracy int) float64 {
	denominator := 1.0
	for i := 0; i < accuracy; i++ {
		denominator = denominator * 10
	}
	randomNumber := float64(utils.NewRandom().Intn(1, 9)) / denominator
	switch mode {
	case 1:
		currentPrice += randomNumber
	case 2:
		currentPrice -= randomNumber
	}
	return currentPrice
}

// 100288.478
