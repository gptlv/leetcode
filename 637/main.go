package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}

	if root == nil {
		return res
	}

	q := []*TreeNode{}
	q = append(q, root)

	for len(q) != 0 {
		s := len(q)
		sum := 0

		for i := 0; i < s; i++ {
			t := q[0]
			if t != nil {
				sum += t.Val
			}
			q = q[1:]

			if t.Left != nil {
				q = append(q, t.Left)
			}

			if t.Right != nil {
				q = append(q, t.Right)
			}
		}

		res = append(res, float64(sum)/float64(s))

	}

	return res

}
