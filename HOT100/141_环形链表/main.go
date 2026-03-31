package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//func hasCycle(head *ListNode) bool {
//	m := make(map[*ListNode]int)
//	var ptr *ListNode
//	ptr = head
//	for ptr != nil {
//		if m[ptr] > 0 {
//			return true
//		}
//		m[ptr]++
//		ptr = ptr.Next
//	}
//	return false
//}
