package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{}
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	head.Next = nil

	return prev
}
