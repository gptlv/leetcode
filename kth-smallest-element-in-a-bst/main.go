package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return k
	}

	res := []int{}

	inOrder(root, &res)
	return res[k-1]
}

func inOrder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inOrder(root.Right, vals)
}
