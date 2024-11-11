package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}

	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length/2; i++ {
			left, right := q[i], q[length-1-i]
			if left == nil && right == nil {
				continue
			}
			if left == nil || right == nil || left.Val != right.Val {
				return false
			}
		}

		newLevel := []*TreeNode{}
		for _, node := range q {
			if node != nil {
				newLevel = append(newLevel, node.Left, node.Right)
			}
		}
		q = newLevel
	}

	return true
}
