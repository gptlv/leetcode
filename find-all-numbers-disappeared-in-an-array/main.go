package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	numsMap := make(map[int]struct{})
	res := []int{}
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	for i := range len(nums) {
		if _, exists := numsMap[i+1]; !exists {
			res = append(res, i+1)
		}
	}

	return res
}
