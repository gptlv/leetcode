package main

import "sort"

func main() {

}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	target := len(nums) / 2
	count := 0

	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	last := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			count++
		} else {
			count = 0
			last = nums[i]
		}

		if count == target {
			res = nums[i]
		}
	}

	return res
}
