package main

import "math"

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)

	for i, num := range nums {
		if _, ok := m[num]; ok && math.Abs(float64(m[num]-i)) <= float64(k) {
			return true
		}

		m[num] = i
	}

	return false

}
