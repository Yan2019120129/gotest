package top4

import (
	"fmt"
	"testing"
)

// ReverseList 反转列表
func TestReverseList(t *testing.T) {
	temp := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	data := ReverseList(temp)
	t.Log(data)
	fmt.Println(data)
}
