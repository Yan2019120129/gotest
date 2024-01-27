package time_test_test

import (
	"gotest/base/time_test"
	"testing"
	"time"
)

// TestTimeCreate 测试time.unix 创建时间是否会不一致
func TestTestTimeCreate(t *testing.T) {
	time_test.TestTimeCreate()
}

// TestTimeTicker 测试时钟
func TestTimeTicker(t *testing.T) {

	time_test.TestTimeTicker(5 * time.Second)
}
