package main

func main() {

}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	k %= len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]

		left++
		right--
	}
}

// func rotate(nums []int, k int) {
// 	if k == 0 {
// 		return
// 	}

// 	k %= len(nums)

// 	left := len(nums) - k
// 	right := len(nums) - 1
// 	target := left - 1

// 	for left > 0 {
// 		moveToTheEnd(nums, target, left, right)
// 		target--
// 		left--
// 		right--
// 	}

// }

// func moveToTheEnd(nums []int, target, left, right int) {
// 	for target != right {
// 		nums[left], nums[target] = nums[target], nums[left]

// 		target++
// 		left++
// 	}
// }
