package swith

import (
	"fmt"
	"time"
)

// Switch 流程测试
func Switch(nanoseconds time.Duration) string {
	Nanoseconds := float64(nanoseconds.Nanoseconds()) / 1e6
	msg := ""
	switch {
	// 大于1秒，小于3秒为黄色
	case nanoseconds > time.Second && nanoseconds < 3*time.Second:
		msg = "黄色"
	// 大于3秒为红色
	case nanoseconds >= 3*time.Second:
		msg = "红色"
	// 默认绿色
	default:
		msg = "绿色"
	}
	fmt.Println(Nanoseconds)
	return fmt.Sprintf("spend:[%v]", msg)
}
