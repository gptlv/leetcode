package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	dummyHead := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			dummy.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			dummy.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}

		dummy = dummy.Next
	}

	if list1 == nil {
		dummy.Next = list2
	}

	if list2 == nil {
		dummy.Next = list1
	}

	return dummyHead.Next
}
