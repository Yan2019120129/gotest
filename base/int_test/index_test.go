package int_test_test

import (
	"gotest/base/int_test"
	"testing"
)

// TestIntSum 测试中间值计算是否会出现超出定义范围
func TestIntSum(t *testing.T) {
	int_test.IntSum()
}

// TestIntSumOverflow 测试中间值计算是否会出现超出定义范围的情况
func TestIntSumOverflow(t *testing.T) {
	int_test.IntSum()
}
