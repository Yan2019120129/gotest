package top4

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var node *ListNode = nil
	for tmp := head; tmp != nil; tmp = tmp.Next {
		newNode := new(ListNode)
		newNode.Val = tmp.Val
		newNode.Next = nil

		// 获取next后的所有节点
		temp := node

		// 将新节点放到头部
		node = newNode

		newNode.Next = temp
	}

	return node
}
