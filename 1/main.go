package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// func twoSum(nums []int, target int) []int {
// 	res := []int{0, 0}
// loop:
// 	for i := 0; i < len(nums); i++ {
// 		res[0] = i
// 		for j := i + 1; j < len(nums); j++ {
// 			if sum := nums[i] + nums[j]; sum == target {
// 				res[1] = j
// 				break loop
// 			}
// 		}
// 	}

// 	return res
// }

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		compliment := target - num
		if idx, found := seen[compliment]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
