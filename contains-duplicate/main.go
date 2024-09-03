package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := numsMap[num]; exists {
			return true
		}
		numsMap[num] = struct{}{}
	}
	return false
}
