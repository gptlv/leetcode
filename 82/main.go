package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	prev := dummy
	cur := head

	for cur != nil {

		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if prev.Next == cur {
			prev = cur //move to the next
		} else {
			prev.Next = cur.Next //skip duplicates, cur points to the last duplicate, hence, we use cur.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}
