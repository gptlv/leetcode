package main

func main() {

}

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	res := make([]int, len(nums))

	left[0] = 1
	right[(len(nums) - 1)] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for j := len(nums) - 2; j >= 0; j-- {
		right[j] = nums[j+1] * right[j+1]
	}

	for k := 0; k < len(nums); k++ {
		res[k] = left[k] * right[k]
	}

	return res
}
