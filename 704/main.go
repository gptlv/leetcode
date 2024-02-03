package main

func main() {

}

func search(nums []int, target int) int {
	lp := 0
	rp := len(nums) - 1
	mid := 0
	for lp <= rp {
		mid = (lp + rp) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			lp = mid + 1

		}

		if nums[mid] > target {
			rp = mid - 1

		}

	}

	return -1
}
