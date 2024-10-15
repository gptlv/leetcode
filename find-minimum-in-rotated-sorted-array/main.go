package main

func main() {

}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if nums[left] <= nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := (left + right) / 2

		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
