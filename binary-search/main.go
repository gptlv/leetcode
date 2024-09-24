package main

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	search(nums, 9)
}

func search(nums []int, target int) int {
	var left, right, middle int
	right = len(nums) - 1

	for left <= right {
		middle = (left + right) / 2

		if nums[middle] < target {
			left = middle + 1
		}

		if nums[middle] > target {
			right = middle - 1
		}

		if nums[middle] == target {
			return middle
		}
	}

	return -1

}
