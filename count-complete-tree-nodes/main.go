package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var count int
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		count++
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return count
}
