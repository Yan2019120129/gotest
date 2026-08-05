package main

import (
	"fmt"
)

func main() {
	// 计算 10^12 * 200ms 的结果，并换算为各种时间单位
	totalSeconds := 1e12 * 0.2 // 200ms = 0.2s
	fmt.Printf("10^12 * 200ms = %.2f 秒\n", totalSeconds)
	fmt.Printf("                 = %.2f 分钟\n", totalSeconds/60)
	fmt.Printf("                 = %.2f 小时\n", totalSeconds/3600)
	fmt.Printf("                 = %.2f 天\n", totalSeconds/86400)
	fmt.Printf("                 = %.2f 月（按30天/月）\n", totalSeconds/86400/30)
	fmt.Printf("                 = %.2f 年（按365天/年）\n", totalSeconds/86400/365)
}
