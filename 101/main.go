package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil || root.Right == nil {
		return root.Left == root.Right
	}

	return isMirror(root.Left, root.Right)

	// root.Right = invertTree(root.Right)
	// return isSame(root.Left, root.Right)
}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return (p.Val == q.Val) && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}

// 	root.Left, root.Right = root.Right, root.Left
// 	root.Left = invertTree(root.Left)
// 	root.Right = invertTree(root.Right)

// 	return root
// }

// func isSame(p *TreeNode, q *TreeNode) bool {
// 	if p == nil || q == nil {
// 		return p == q
// 	}

// 	return (p.Val == q.Val) && isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
// }
