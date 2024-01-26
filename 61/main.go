package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	length := getLength(head)

	if length == 0 || length == 1 {
		return head
	}

	k = k % length

	return rotate(head, k)
}

func getLength(head *ListNode) int {
	length := 0
	cur := head

	for cur != nil {
		cur = cur.Next
		length++
	}

	return length
}

func rotate(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	cur := head

	for cur.Next.Next != nil {
		cur = cur.Next
	}

	last := cur.Next
	last.Next = head

	cur.Next = nil
	k--

	return rotate(last, k)
}
