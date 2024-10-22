package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	helper(root, &res)

	return res
}

func helper(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, diameter)
	right := helper(root.Right, diameter)

	*diameter = max(*diameter, left+right)

	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
