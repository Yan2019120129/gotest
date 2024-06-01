package top2

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

// TestMoveZeroes 测试移动零值
func TestMoveZeroes(t *testing.T) {
	baseArray := []int{2, 0, 4, 5, 0, 6, 0, 0, 6}
	MoveZeroes(baseArray)
	t.Log(baseArray)
}

// BenchmarkMoveZeroes 测试移动零值
func BenchmarkMoveZeroes(b *testing.B) {
	//baseArray := []int{2, 0, 4, 5, 0, 6, 0, 0, 6}
	baseArray := make([]int, 0)
	for i := 0; i < 10; i++ {
		baseArray = append(baseArray, gofakeit.Number(0, 100))
	}
	//b.Log(baseArray)
	for i := 0; i < b.N; i++ {
		MoveZeroes(baseArray)
	}
	b.Log(baseArray)
}

// BenchmarkMoveZeroesMy 测试移动零值
func BenchmarkMoveZeroesMy(b *testing.B) {
	//baseArray := []int{2, 0, 4, 5, 0, 6, 0, 0, 6}
	baseArray := make([]int, 0)
	for i := 0; i < 10; i++ {
		baseArray = append(baseArray, gofakeit.Number(0, 100))
	}
	//b.Log(baseArray)
	for i := 0; i < b.N; i++ {
		MoveZeroesMy(baseArray)
	}
	//b.Log(baseArray)
}
