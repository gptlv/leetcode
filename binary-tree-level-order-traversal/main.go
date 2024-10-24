package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}
	res := [][]int{}

	if root == nil {
		return res
	}

	for len(q) > 0 {
		currentLength := len(q)
		level := []int{}

		for i := 0; i < currentLength; i++ {
			node := q[i]
			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
		q = q[currentLength:]
	}

	return res
}
