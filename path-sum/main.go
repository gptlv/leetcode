package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	sum := 0
	return helper(root, sum, targetSum)
}

func helper(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val

	if sum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return helper(root.Left, sum, targetSum) || helper(root.Right, sum, targetSum)
}
