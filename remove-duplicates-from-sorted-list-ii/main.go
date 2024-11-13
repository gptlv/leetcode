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

	dummy := &ListNode{Next: head}
	distinct := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur = skipDuplicates(cur)
			distinct.Next = cur
		} else {
			distinct = cur
			cur = cur.Next
		}

	}

	return dummy.Next
}

func skipDuplicates(head *ListNode) *ListNode {
	cur := head
	val := cur.Val

	for cur != nil && cur.Val == val {
		cur = cur.Next
	}

	return cur
}
