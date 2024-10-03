package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	ptr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[ptr] = nums[i]
			ptr++
		}
	}

	return ptr
}
