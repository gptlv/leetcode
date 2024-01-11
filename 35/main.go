package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 13, 16, 89}
	fmt.Println(searchInsert(nums, 10))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		v := nums[m]
		if target < v {
			r = m - 1
		}
		if target > v {
			l = m + 1
		}
		if target == v {
			return m
		}
	}
	return l
}
