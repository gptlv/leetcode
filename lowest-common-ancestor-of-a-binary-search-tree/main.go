package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return check(root, p, q)
}

func check(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return check(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return check(root.Right, p, q)
	}

	return root
}
