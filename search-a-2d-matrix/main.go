package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	return outerBinarySearch(matrix, target)
}

func innerBinarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func outerBinarySearch(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1

	for top <= bottom {
		mid := (top + bottom) / 2
		if target > matrix[mid][len(matrix[mid])-1] {
			top = mid + 1
		} else if target < matrix[mid][0] {
			bottom = mid - 1
		} else {
			return innerBinarySearch(matrix[mid], target)
		}
	}

	return false
}
