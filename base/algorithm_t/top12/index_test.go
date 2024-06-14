package top12

import (
	"fmt"
	"testing"
)

var GroupAnagramsData = [][]string{
	{"eat", "tea", "tan", "ate", "nat", "bat"},
	{""},
	{"a"},
}

// TestGroupAnagrams 测试异位分词分组
func TestGroupAnagrams(t *testing.T) {
	for _, datum := range GroupAnagramsData {
		fmt.Println(groupAnagrams(datum))
	}
}
