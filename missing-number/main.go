package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
}

func missingNumber(nums []int) int {
	numsMap := make(map[int]struct{})

	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	for i := range len(nums) {
		if _, exists := numsMap[i]; !exists {
			return i
		}
	}

	return len(nums)
}
