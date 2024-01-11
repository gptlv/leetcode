package main

func main() {

	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

}

func removeDuplicates(nums []int) int {
	idx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			nums[idx+1] = nums[i]
			idx++
		}
	}
	return idx + 1
}
