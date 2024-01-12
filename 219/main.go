package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	idx := 0

	for k != 0 {
		if idx+k >= len(nums) {
			k--
			idx = 0
			continue
		}

		if nums[idx] == nums[idx+k] {
			return true
		}

		idx++
	}

	return false
}
