package algorithm

import (
	"fmt"
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
