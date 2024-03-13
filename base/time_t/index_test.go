package time_t

import (
	"testing"
	"time"
)

// TimeCreate 测试time.unix 创建时间是否会不一致
func TestTestTimeCreate(t *testing.T) {
	TimeCreate()
}

// TestTimeTicker 测试时钟
func TestTimeTicker(t *testing.T) {

	TimeTicker(5 * time.Second)
}
