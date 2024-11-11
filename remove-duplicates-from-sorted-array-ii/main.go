package main

func main() {

}

func removeDuplicates(nums []int) int {
	idx := 0

	for i := 0; i < len(nums); i++ {
		if idx < 2 || nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
