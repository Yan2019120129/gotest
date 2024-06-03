package top3

import (
	"testing"
)

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
