package main

func main() {

}

func subsets(nums []int) [][]int {
	var res [][]int
	n := len(nums)

	for i := 0; i < (1 << n); i++ {
		var subset []int

		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				subset = append(subset, nums[j])
			}

		}
		res = append(res, subset)
	}

	return res
}
