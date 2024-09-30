package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil && tree2 != nil || tree1 != nil && tree2 == nil {
		return false
	}

	return tree1.Val == tree2.Val && isSame(tree1.Left, tree2.Left) && isSame(tree1.Right, tree2.Right)
}
