package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}
