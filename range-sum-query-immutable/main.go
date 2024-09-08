package main

import (
	"fmt"
)

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
}

// type NumArray struct {
// 	data map[int]int
// }

// func Constructor(nums []int) NumArray {
// 	data := make(map[int]int, len(nums))
// 	for i, num := range nums {
// 		data[i] = num
// 	}

// 	return NumArray{data: data}

// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	sum := 0
// 	for i := left; i <= right; i++ {
// 		sum += this.data[i]
// 	}

// 	return sum
// }

type NumArray struct {
	nums  []int
	cache []int
}

func Constructor(nums []int) NumArray {
	cache := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		cache[i] = sum
	}

	return NumArray{nums: nums, cache: cache}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.cache[right] - this.cache[left] + this.nums[left]
}
