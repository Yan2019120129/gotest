package string_test_test

import (
	"fmt"
	"gotest/base/string_test"
	"testing"
)

// TestFindValue 查找值
func TestFindValue(t *testing.T) {
	string_test.FindValue()
}

// TestScan 测试控制台输入参数
func TestScan(t *testing.T) {
	message := ""
	fmt.Print("等待输入：")
	_, err := fmt.Scan(&message)
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("输入内容:", message)
}

// TestString 测试[]string 转换[]byte
func TestString(t *testing.T) {

}

// TestStrconvFromInt 测试字符串转换十进制
func TestStrconvFromInt(t *testing.T) {
	string_test.TestStrconv()
}

// TestMultipleDataString 测试传入多数据的情况
func TestMultipleDataString(t *testing.T) {
	string_test.MultipleDataString("yan", "jia", "jie")
}

// TestMultipleDataInt 测试传入多数据的情况
func TestMultipleDataInt(t *testing.T) {
	//string_test.MultipleDataInt(45, 33, 90)
	string_test.MultipleDataInt(45)
}
