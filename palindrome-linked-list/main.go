package main

import "fmt"

func main() {
	node4 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Println(isPalindrome(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	cur := makeCopy(head)
	curReversed := reverseList(head)

	for cur != nil {
		fmt.Println(cur.Val, curReversed.Val)
		if cur.Val != curReversed.Val {
			return false
		}

		cur = cur.Next
		curReversed = curReversed.Next
	}
	return true
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

func makeCopy(head *ListNode) *ListNode {
	dummy := &ListNode{}
	copy := dummy
	cur := head

	for cur != nil {
		t := &ListNode{Val: cur.Val}
		copy.Next = t
		copy = copy.Next
		cur = cur.Next
	}

	return dummy.Next
}
