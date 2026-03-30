package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, prev *ListNode
	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func endfirsthalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next.Next != nil && slow.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	tail := endfirsthalf(head)
	p2 := reverseList(tail.Next)
	p1 := head
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	tail.Next = reverseList(p2)
	return result
}
