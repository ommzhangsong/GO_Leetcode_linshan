package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	ptr := head

	for ptr != nil {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}
	return prev
}
