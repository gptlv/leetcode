package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		currentLength := len(q)

		for i := 0; i < currentLength; i++ {
			node := q[i]

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			if i == currentLength-1 {
				res = append(res, node.Val)
			}
		}

		q = q[currentLength:]
	}

	return res
}
