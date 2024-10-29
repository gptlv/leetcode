package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := root.Val
	helper(root, &res)

	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, res)
	if left < 0 {
		left = 0
	}

	right := helper(root.Right, res)
	if right < 0 {
		right = 0
	}

	*res = max(*res, root.Val+left+right)

	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
