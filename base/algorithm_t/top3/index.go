package top3

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func searchValues(headB *ListNode) ([]string, int, map[string]*ListNode) {
	mapNode := make(map[string]*ListNode)
	values := make([]string, 0)
	for headB != nil {
		key := fmt.Sprintf("%p", headB)
		mapNode[key] = headB
		values = append(values, key)
		headB = headB.Next
	}

	return values, len(values), mapNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	values, lenght, tempMap := searchValues(headA)
	values1, lenght1, _ := searchValues(headB)
	tempLen := lenght
	if lenght > lenght1 {
		tempLen = lenght1
	}
	index := ""
	for i := 1; i <= tempLen; i++ {
		if values1[lenght1-i] == values[lenght-i] {
			index = values1[lenght1-i]
			continue
		}
	}
	return tempMap[index]
}

func GetIntersectionNodeTow(headA, headB *ListNode) *ListNode {
	mapNode := make(map[*ListNode]bool)
	for {
		if headA != nil {
			if _, ok := mapNode[headA]; ok {
				return headA
			}
			mapNode[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := mapNode[headB]; ok {
				return headB
			}
			mapNode[headB] = true
			headB = headB.Next
		}
		if headA == nil && headB == nil {
			return nil
		}
	}
}
