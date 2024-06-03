package top10

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 交换链表节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := head.Next
	head.Next = swapPairs(newNode.Next)
	newNode.Next = head
	return newNode
}

// swapPairs 交换链表节点
func swapPairsTwo(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1
	}
	return dummyHead.Next
}
