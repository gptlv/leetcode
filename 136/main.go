package main

import "sort"

func main() {

}

func singleNumber(nums []int) int {
	length := len(nums)

	if length == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	if nums[0] != nums[1] {
		return nums[0]
	}

	for i := 1; i < length-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[length-1]

}
