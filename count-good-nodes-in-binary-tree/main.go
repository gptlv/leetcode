package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	res := 0

	if root == nil {
		return res
	}

	helper(root, root.Val, &res)

	return res
}

func helper(root *TreeNode, max int, res *int) {
	if root == nil {
		return
	}

	if root.Val >= max {
		*res++
		max = root.Val
	}

	helper(root.Left, max, res)
	helper(root.Right, max, res)
}
