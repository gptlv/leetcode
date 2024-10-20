package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := getMid(head)
	secondHalf := reverse(mid.Next)
	mid.Next = nil
	mergeLists(head, secondHalf)
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head

	for cur != nil {
		t := cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}

	return prev
}

func mergeLists(left, right *ListNode) {
	for left != nil && right != nil {
		temp1 := left.Next
		temp2 := right.Next

		left.Next = right
		right.Next = temp1

		left = temp1
		right = temp2
	}
}
