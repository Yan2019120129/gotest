package algorithm

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

// TestTwoSum 两数之和
func TestTwoSum(t *testing.T) {
	data := TwoSum([]int{2, 7, 15, 11}, 9)
	fmt.Println("data:", data)
}
func TestClimbStairsRecursionWithOneStep(t *testing.T) {
	result := ClimbStairsRecursion(1)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestClimbStairsRecursionWithTwoSteps(t *testing.T) {
	result := ClimbStairsRecursion(2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestClimbStairsRecursionWithThreeSteps(t *testing.T) {
	result := ClimbStairsRecursion(3)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestClimbStairsRecursionWithTenSteps(t *testing.T) {
	result := ClimbStairsRecursion(10)
	if result != 89 {
		t.Errorf("Expected 89, got %d", result)
	}
}

func TestClimbStairsRecursionWithNegativeSteps(t *testing.T) {
	result := ClimbStairsRecursion(-1)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// TestMaximumNumberOfStringPairs 获取取求反相同的字符串总数
func TestMaximumNumberOfStringPairs(t *testing.T) {
	//count := MaximumNumberOfStringPairs([]string{"ab", "ba", "cc"})
	//count := MaximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"})
	count := MaximumNumberOfStringPairs([]string{"aa", "ab"})
	fmt.Println("count:", count)
}

// TestMaximumNumberOfStringPairsTwo 获取取求反相同的字符串总数
func TestMaximumNumberOfStringPairsTwo(t *testing.T) {
	count := MaximumNumberOfStringPairsTow([]string{"ab", "ba", "cc"})
	//count := MaximumNumberOfStringPairsTow([]string{"cd", "ac", "dc", "ca", "zz"})
	//count := MaximumNumberOfStringPairsTow([]string{"aa", "ab"})
	fmt.Println("count:", count)
}

// TestMaximumNumberOfStringPairsFrid 获取取求反相同的字符串总数
func TestMaximumNumberOfStringPairsFrid(t *testing.T) {
	count := MaximumNumberOfStringPairsFrid([]string{"ab", "ba", "cc"})
	//count := MaximumNumberOfStringPairsFrid([]string{"cd", "ac", "dc", "ca", "zz"})
	//count := MaximumNumberOfStringPairsFrid([]string{"aa", "ab"})
	fmt.Println("count:", count)
}

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

// TestGetIntersectionNode 测试交叉链表
func TestGetIntersectionNode(t *testing.T) {
	//sync := &ListNode{
	//	Val: 8,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val: 5},
	//	},
	//}
	//temp := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: sync}}
	//temp2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: sync}}}
	sync := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	temp2 := &ListNode{Val: 3, Next: sync}
	temp := &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: sync}}}

	//sync := &ListNode{
	//	Val: 1,
	//}
	//temp := sync
	//temp2 := sync

	node := GetIntersectionNodeTow(temp, temp2)
	if node != nil {
		t.Log(node.Val)
	}
}

// BenchmarkGetIntersectionNode 测试交叉链表
func BenchmarkGetIntersectionNode(b *testing.B) {
	temp := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	temp2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}
	for i := 0; i < b.N; i++ {
		b.Log(GetIntersectionNode(temp, temp2))
	}
}
