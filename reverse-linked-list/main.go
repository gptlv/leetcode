package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head.Next
	last := head
	prev := head

	for cur != nil {
		t := cur
		cur = cur.Next
		t.Next = prev
		prev = t
	}

	last.Next = nil

	return prev
}
