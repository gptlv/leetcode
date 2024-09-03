package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	numsFrequency := make(map[int]int)

	for _, num := range nums {
		numsFrequency[num] = numsFrequency[num] + 1
	}

	for key, value := range numsFrequency {
		if value != 2 {
			return key
		}
	}

	return nums[0]
}
