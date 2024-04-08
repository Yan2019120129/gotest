package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	number := 4
	x := 1.0000000
	sum := 100.0
	for i := 0; i < number; i++ {
		x = x * 10
	}
	sum = sum * 1 / x
	b := randNum(1, 10)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(sum)
}

func randNum(m, n int) int {
	// 设置随机数种子
	rand.NewSource(time.Now().UnixNano())

	// 生成大于m且小于n的随机整数
	return rand.Intn(n-m) + m
}
