package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	helper(root)
}

var prev *TreeNode

func helper(root *TreeNode) {
	if root == nil {
		return
	}

	l, r := root.Left, root.Right
	root.Left, root.Right = nil, nil

	if prev != nil {
		prev.Right = root
	}

	prev = root

	helper(l)
	helper(r)
}
