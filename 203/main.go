package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	last := &ListNode{Next: head}
	res := last

	for cur != nil {
		if cur.Val != val {
			last.Next = cur
			last = last.Next
		}
		cur = cur.Next
	}

	if last.Next != nil && last.Next.Val == val {
		last.Next = nil
	}

	return res.Next
}
