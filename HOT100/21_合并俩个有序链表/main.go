package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := new(ListNode)
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return cur.Next
}
