package top8

type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome 是否回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeList := make([]*ListNode, 0)
	for tmp := head; tmp != nil; tmp = tmp.Next {
		nodeList = append(nodeList, tmp)
	}
	nodeLen := len(nodeList)
	for i := 0; i < nodeLen/2; i++ {
		if nodeList[i].Val != nodeList[nodeLen-1-i].Val {
			return false
		}
	}

	return true
}

// isPalindromeTow 是否回文链表
func isPalindromeTow(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := findMid(head)
	mid = mid.Next
	mid = reverseList(mid)
	for ; mid != nil; mid = mid.Next {
		if mid.Val != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}

// reverseList 撤销链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for Next := head.Next; head != nil; Next = Next.Next {
		head.Next = pre
		pre = head
		head = Next
		if Next == nil {
			break
		}
	}
	return pre
}

// findMid 查找中间节点
func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
