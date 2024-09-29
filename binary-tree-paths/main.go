package main

import (
	"fmt"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := &[]string{}

	helper(root, "", res)

	return *res
}

func helper(root *TreeNode, path string, res *[]string) {
	var newPath string
	if path == "" {
		newPath = strconv.Itoa(root.Val)
	} else {
		newPath = fmt.Sprintf("%s->%v", path, root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, newPath)
	}

	if root.Left != nil {
		helper(root.Left, newPath, res)
	}

	if root.Right != nil {
		helper(root.Right, newPath, res)
	}
}
