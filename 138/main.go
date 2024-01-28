package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	m := make(map[*Node]*Node)

	cur := head

	dummy := &Node{}
	dummyHead := dummy

	for cur != nil {
		dummy.Next = &Node{Val: cur.Val, Random: cur.Random}
		dummy = dummy.Next

		m[cur] = dummy

		cur = cur.Next

	}

	cur = dummyHead.Next

	for cur != nil {
		if _, ok := m[cur.Random]; ok {
			cur.Random = m[cur.Random]
		}

		cur = cur.Next
	}

	return dummyHead.Next

}
