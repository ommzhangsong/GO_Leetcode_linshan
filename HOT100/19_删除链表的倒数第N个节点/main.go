package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head
	count := 0
	for fast != nil {
		if count == n {
			slow.Next = slow.Next.Next
		}
		count++
		fast = fast.Next
		slow = slow.Next
	}
	return dummy.Next
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil {
//		return nil
//	}
//	dummy := &ListNode{Next: head}
//	count := 0
//	ptr := head
//	for ptr != nil {
//		count++
//		ptr = ptr.Next
//	}
//	cur := dummy
//	for i := 0; i < count-n; i++ {
//		cur = cur.Next
//	}
//	cur.Next = cur.Next.Next
//	return dummy.Next
//}
