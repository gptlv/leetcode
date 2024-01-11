package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 6}

	node1.Next = node3
	node3.Next = node4
	node4.Next = node6

	fmt.Println("First linked list:")
	printList(node1)

	// Second linked list: 2 -> 3 -> 5 -> 7 -> nil
	node2 := &ListNode{Val: 2}
	node3_2 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 7}

	node2.Next = node3_2
	node3_2.Next = node5
	node5.Next = node7

}

func printList(list *ListNode) {
	current := &list
	for current != nil {
		fmt.Println(current.Val)
		current = &current.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

}
