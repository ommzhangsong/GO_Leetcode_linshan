package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	ptr := l3
	i := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + i
		cur := sum % 10
		i = sum / 10
		l3.Next = &ListNode{Val: cur}
		l3 = l3.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil && i >= 1 {
		l3.Val = 1
	}
	if l1 == nil {
		for l2 != nil {
			sum := l2.Val + i
			cur := sum % 10
			i = sum / 10
			l3.Next = &ListNode{Val: cur}
			l3 = l3.Next
			l2 = l2.Next
		}
	} else {
		for l1 != nil {
			sum := l1.Val + i
			cur := sum % 10
			i = sum / 10
			l3.Next = &ListNode{Val: cur}
			l3 = l3.Next
			l1 = l1.Next
		}
	}
	return ptr.Next
}
