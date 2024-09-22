package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	if cur == nil {
		return cur
	}

	new := &ListNode{Val: head.Val}
	newHead := new
	max := cur.Val

	for cur != nil {
		if cur.Val > max {
			new.Next = &ListNode{Val: cur.Val}
			new = new.Next
		}

		max = cur.Val
		cur = cur.Next
	}

	return newHead
}
