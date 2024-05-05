package string_t

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// TestFindValue 查找值
func TestFindValue(t *testing.T) {
	FindValue()
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

func BenchmarkTypeConversion(b *testing.B) {
	byteSlice := []byte("Hello")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(byteSlice)
	}
}

func BenchmarkStrconvFormatInt(b *testing.B) {
	byteSlice := []byte("Hello")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(byteSlice[i%len(byteSlice)]), 10)
	}
}
func TestStrconvFormatInt(t *testing.T) {
	fmt.Println("BenchmarkTypeConversion:")
	fmt.Println(testing.Benchmark(BenchmarkTypeConversion))

	fmt.Println("\nBenchmarkStrconvFormatInt:")
	fmt.Println(testing.Benchmark(BenchmarkStrconvFormatInt))
}

// TestStrconvFromInt 测试字符串转换十进制
func TestStrconvFromInt(t *testing.T) {
	TestStrconv()
}

// TestMultipleDataString 测试传入多数据的情况
func TestMultipleDataString(t *testing.T) {
	MultipleDataString("yan", "jia", "jie")
}

// TestMultipleDataInt 测试传入多数据的情况
func TestMultipleDataInt(t *testing.T) {
	//MultipleDataInt(45, 33, 90)
	MultipleDataInt(45)
}

// TestMultipleDataInt 测试传入多数据的情况
func TestStringSlice(t *testing.T) {
	stringList := []string{"yan"}
	joinStrings := strings.Join(stringList, ",")
	fmt.Println(joinStrings)
}
