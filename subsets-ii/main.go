package main

import "sort"

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var dfs func(i int, subset []int)

	dfs = func(i int, subset []int) {
		if i == len(nums) {
			c := make([]int, len(subset))
			copy(c, subset)
			res = append(res, c)
			return
		}

		subset = append(subset, nums[i])
		dfs(i+1, subset)

		subset = subset[:len(subset)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

		dfs(i+1, subset)

	}

	dfs(0, []int{})

	return res
}
