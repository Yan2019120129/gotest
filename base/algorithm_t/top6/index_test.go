package top6

import (
	"fmt"
	"testing"
)

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
