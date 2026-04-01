package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := prev.Next.Next

		a.Next = b.Next
		b.Next = a
		prev.Next = b

		prev = a
	}

	return dummy.Next
}
