package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	ptr := dummy

	for i := 0; i <= n; i++ {
		ptr = ptr.Next
	}

	for ptr != nil {
		cur = cur.Next
		ptr = ptr.Next
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}
