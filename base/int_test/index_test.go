package int_test_test

import (
	"fmt"
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

// TestOnMessage 测试方法类型的使用
func TestOnMessage(t *testing.T) {
	int_test.Instance.OnMessage = SetMessage
	int_test.Instance.ForMessage("zhe", "ge", "fang", "fa", "ke", "yi", "de")
}

func SetMessage(msg string) string {
	fmt.Println(msg)
	return msg
}
