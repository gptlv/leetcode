package main

func main() {

}

func removeElement(nums []int, val int) int {
	ptr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ptr] = nums[i]
			ptr++
		}
	}

	return ptr
}
