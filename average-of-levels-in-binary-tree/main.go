package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	res := []float64{}

	for len(queue) > 0 {
		length := len(queue)
		sum := 0

		for i := 0; i < length; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			sum += queue[0].Val
			queue = queue[1:]

		}

		res = append(res, float64(sum)/float64(length))
	}

	return res
}
