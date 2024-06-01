package top7

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle 是否环形链表
func HasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if _, ok := nodeMap[tmp]; ok {
			return true
		}
		nodeMap[tmp] = false
	}
	return false
}

// HasCycleTwo 是否环形链表2
func HasCycleTwo(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]*ListNode)
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if p, ok := nodeMap[tmp]; ok {
			return p
		}
		nodeMap[tmp] = tmp
	}
	return nil
}
