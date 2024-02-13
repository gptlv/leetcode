package main

func main() {

}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	res := 0

	helper(sum, nums, target, &res)

	return res

}

func helper(sum int, nums []int, target int, res *int) {
	if len(nums) == 0 {
		if sum == target {
			*res++
		}

		return
	}

	helper(sum+nums[0], nums[1:], target, res)
	helper(sum-nums[0], nums[1:], target, res)
}
