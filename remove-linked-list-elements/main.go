package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	dummy := &ListNode{}
	dummyHead := dummy

	for cur != nil {
		if cur.Val != val {
			t := &ListNode{Val: cur.Val}
			dummy.Next = t
			dummy = dummy.Next
		}

		cur = cur.Next
	}

	return dummyHead.Next
}
