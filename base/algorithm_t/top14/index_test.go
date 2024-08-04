package top14

import (
	"fmt"
	"testing"
)

// 4 9 25 49 81 121 169
func TestNonSpecialCount(t *testing.T) {
	testCase := [][]int{
		{1, 4},
		{1, 2},
		{4, 5},
		{5, 7},
		{4, 16},
		{182, 18677},
		{28655635, 337909353},
	}
	for _, ints := range testCase {
		l := ints[0]
		r := ints[1]
		val := nonSpecialCount(l, r)
		fmt.Println(val, r-l+1)
	}
}
func BenchmarkSieveOfEratosthenes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
