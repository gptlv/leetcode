package main

func main() {

}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}

	num = num*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return num
	}

	return helper(root.Left, num) + helper(root.Right, num)
}
