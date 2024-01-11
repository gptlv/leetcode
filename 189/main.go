package main

import "fmt"

// func rotate(nums []int, k int) {
// 	length := len(nums)
// 	if k > length {
// 		k = k % length
// 	}
// 	nums = append(nums[length-k:], nums[:length-k]...)
// 	fmt.Println(nums)
// }

func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := make([]int, len(nums))
	copy(temp, nums)
	for i := 0; i < len(nums); i++ {
		j := (i + k) % len(nums)
		nums[j] = temp[i]
	}

}

// func rotate(nums []int, k int) {
// 	k = k % len(nums)
// 	temp := make([]int, len(nums))
// 	copy(temp, nums)

// 	for i := 0; i < len(nums); i++ {
// 		j := (i + k) % len(nums)
// 		nums[j] = temp[i]
// 	}
// }

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

}
