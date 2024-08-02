package top1

import (
	"fmt"
	"strconv"
	"testing"
)

// TestTwoSum 两数之和
func TestTwoSum(t *testing.T) {
	data := TwoSum([]int{2, 7, 15, 11}, 9)
	fmt.Println("data:", data)
}

//// Benchmark_test 基准测试
//func Benchmark_test(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		tmp := make([]int, 0)
//		for j := 0; j < 100; j++ {
//			tmp = append(tmp, int(gofakeit.Int8()))
//		}
//		sum := int(gofakeit.Int8())
//		//data := TwoSum(tmp, sum)
//		_ = TwoSum(tmp, sum)
//		//firstIndex := 0
//		//endIndex := 0
//		//if len(data) > 0 {
//		//	firstIndex = data[0]
//		//	endIndex = data[1]
//		//}
//		//fmt.Println("data:", data, sum, tmp[firstIndex], tmp[endIndex])
//	}
//}

// Sprintf
func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

// Format
func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

// Itoa
func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
