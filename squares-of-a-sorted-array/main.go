package main

func main() {

}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	left := 0
	right := len(nums) - 1
	i := len(nums) - 1
	for i >= 0 {
		if abs(nums[left]) > abs(nums[right]) {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
		i--
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
