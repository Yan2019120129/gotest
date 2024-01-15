package algorithm_test_test

import (
	"fmt"
	"gotest/base/algorithm_test"
	"testing"
)

// TestTwoSum 两数之和
func TestTwoSum(t *testing.T) {
	data := algorithm_test.TwoSum([]int{2, 7, 15, 11}, 9)
	fmt.Println("data:", data)
}
func TestClimbStairsRecursionWithOneStep(t *testing.T) {
	result := algorithm_test.ClimbStairsRecursion(1)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestClimbStairsRecursionWithTwoSteps(t *testing.T) {
	result := algorithm_test.ClimbStairsRecursion(2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestClimbStairsRecursionWithThreeSteps(t *testing.T) {
	result := algorithm_test.ClimbStairsRecursion(3)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestClimbStairsRecursionWithTenSteps(t *testing.T) {
	result := algorithm_test.ClimbStairsRecursion(10)
	if result != 89 {
		t.Errorf("Expected 89, got %d", result)
	}
}

func TestClimbStairsRecursionWithNegativeSteps(t *testing.T) {
	result := algorithm_test.ClimbStairsRecursion(-1)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
