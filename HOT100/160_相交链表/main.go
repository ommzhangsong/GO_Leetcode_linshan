package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var ptra *ListNode = headA
	var ptrb *ListNode = headB
	for ptra != ptrb {
		if ptra.Next != nil {
			ptra = ptra.Next
		} else {
			ptra = headB
		}
		if ptrb.Next != nil {
			ptrb = ptrb.Next
		} else {
			ptrb = headA
		}
	}
	return ptra
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	m := make(map[*ListNode]bool)
//	var ptra *ListNode = headA
//	var ptrb *ListNode = headB
//	for ptra != nil {
//		m[ptra] = true
//		ptra = ptra.Next
//	}
//	for ptrb != nil {
//		if m[ptrb] {
//			return ptrb
//		}
//		ptrb = ptrb.Next
//	}
//
//	return nil
//
//}
