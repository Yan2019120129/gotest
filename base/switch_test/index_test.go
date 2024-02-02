package swith_test

import (
	"fmt"
	swith "gotest/base/switch_test"
	"testing"
	"time"
)

// TestSwitch 测试流程
func TestSwitch(t *testing.T) {
	elapsed := time.Since(time.Time{})
	fmt.Println(3 * time.Second)
	fmt.Println(swith.Switch(elapsed))
}
