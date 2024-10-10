package main

func main() {

}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for right > left {
		h, w := 0, right-left

		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		s := h * w
		if s > max {
			max = s
		}
	}

	return max
}

// func maxArea(height []int) int {
// 	areaLeft := maxAreaLeft(height)
// 	areaRight := maxAreaRight(height)

// 	if areaLeft > areaRight {
// 		return areaLeft
// 	}

// 	return areaRight
// }

// func maxAreaRight(height []int) int {
// 	right := len(height) - 1
// 	left := len(height) - 2

// 	max := 0

// 	for right >= 1 {
// 		for left >= 0 {
// 			if height[left] >= height[right] {
// 				s := (right - left) * height[right]
// 				if s > max {
// 					max = s
// 				}
// 			}
// 			left--
// 		}
// 		right--
// 		left = right
// 	}

// 	return max
// }

// func maxAreaLeft(height []int) int {
// 	left := 0
// 	right := 1

// 	max := 0

// 	for left < len(height) {
// 		for right < len(height) {
// 			if height[right] >= height[left] {
// 				s := (right - left) * height[left]
// 				if s > max {
// 					max = s
// 				}
// 			}
// 			right++
// 		}
// 		left++
// 		right = left
// 	}

// 	return max
// }
