package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		target := -nums[k]
		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[i] + nums[j]
			if sum > target {
				j--
			} else if sum < target {
				i++
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}

				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			}
		}
	}

	return res
}
