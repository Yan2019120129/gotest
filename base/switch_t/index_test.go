package swith_t

import (
	"fmt"
	"testing"
	"time"
)

// TestSwitch 测试流程
func TestSwitch(t *testing.T) {
	elapsed := time.Since(time.Time{})
	fmt.Println(3 * time.Second)
	fmt.Println(Switch(elapsed))
}

// TestSwitchBool 测试bool
func TestSwitchBool(t *testing.T) {
	month := 3
	switch month {
	case 3, 4, 5:
		fmt.Printf("春天")
		fallthrough
	case 6, 7, 8:
		fmt.Printf("夏天")
		//fallthrough
	case 9, 10, 11:
		fmt.Printf("秋天")
		fallthrough
	case 12, 1, 2:
		fmt.Printf("冬天")
	}
}
