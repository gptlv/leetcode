package main

func removeElement(nums []int, val int) int {
	idx := 0

	for _, num := range nums {
		if nums != val {
			nums[idx] = num
			idx++
		}
	}
	return idx
}
func main() {
	removeElement([]int{2, 2, 3}, 2)
}
