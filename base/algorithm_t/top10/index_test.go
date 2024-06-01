package top10

import "testing"

// TestSwapPairs 两两交换链表中的节点
func TestSwapPairs(t *testing.T) {
	//temp := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	temp := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	//data := isPalindrome(temp)
	//data := swapPairs(temp)
	data := swapPairsTwo(temp)
	t.Log(data)
}
