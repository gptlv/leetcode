package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	less := &ListNode{0, head}
	lHead := less

	greater := &ListNode{0, head}
	gHead := greater

	cur := head

	for cur != nil {
		if cur.Val < x {
			less.Next = cur
			less = less.Next
		} else {
			greater.Next = cur
			greater = greater.Next
		}

		cur = cur.Next
	}

	greater.Next = nil

	less.Next = gHead.Next

	return lHead.Next

}
