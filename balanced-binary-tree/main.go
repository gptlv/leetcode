package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := height(root.Left)
	rightH := height(root.Right)

	if leftH == -1 || rightH == -1 {
		return -1
	}

	if abs(leftH-rightH) > 1 {
		return -1
	}

	return 1 + max(leftH, rightH)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
