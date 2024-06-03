package top8

import "testing"

// TestIsPalindrome 回文列表
func TestIsPalindrome(t *testing.T) {
	temp := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	//temp := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	//data := isPalindrome(temp)
	data := isPalindromeTow(temp)
	t.Log(data)
}
