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
