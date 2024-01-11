package main

import "fmt"

func majorityElement(nums []int) int {
	length := len(nums)
	var element int
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	fmt.Println(freq)

	for key, value := range freq {
		if value > length/2 {
			element = key
			break
		}
	}

	return element
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
