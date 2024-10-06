package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, num := range nums {
		if seen, ok := m[num]; ok {
			if i-seen <= k {
				return true
			}
		}
		m[num] = i
	}

	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	if k == 0 {
// 		return false
// 	}

// 	left, right := 0, 0
// 	delta := 1

// 	for delta <= k {
// 		right = left + delta
// 		if right >= len(nums) {
// 			left = 0
// 			right = 0
// 			delta++
// 			continue
// 		} else if nums[left] == nums[right] {
// 			return true
// 		}

// 		left++
// 		right++
// 	}

// 	return false
// }
