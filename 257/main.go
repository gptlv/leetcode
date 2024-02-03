package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res = make([]string, 0)
	helper(root, "", &res)
	return res
}

func helper(root *TreeNode, base string, res *[]string) {
	if root == nil {
		return
	}

	if base != "" {
		base += "->"
	}

	base += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, base)
	} else {
		helper(root.Left, base, res)
		helper(root.Right, base, res)
	}
}
