package main

import "sort"

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	var dfs func(i int, nums []int, sum int)
	sort.Ints(candidates)
	res := [][]int{}

	dfs = func(i int, nums []int, sum int) {
		if sum == target {
			c := make([]int, len(nums))
			copy(c, nums)
			res = append(res, c)
			return
		}

		if sum > target || i >= len(candidates) {
			return
		}

		nums = append(nums, candidates[i])
		dfs(i+1, nums, sum+candidates[i])

		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}

		nums = nums[:len(nums)-1]
		dfs(i+1, nums, sum)
	}

	dfs(0, []int{}, 0)

	return res
}
