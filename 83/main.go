package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for {
		if cur == nil || cur.Next == nil {
			break
		}

		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			// cur = cur.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// inital solution
//
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	res := &ListNode{Val: head.Val}
// 	resHead := res

// 	val := res.Val

// 	for current := head.Next; current != nil; current = current.Next {

// 		if val != current.Val {
// 			val = current.Val
// 			res.Next = &ListNode{Val: val}
// 			res = res.Next
// 		}

// 	}

// 	return resHead
// }
