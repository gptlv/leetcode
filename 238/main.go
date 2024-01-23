package main

func main() {

	productExceptSelf([]int{4, 10, 8, 2, 9, 6, 5, 11})

}

func productExceptSelf(nums []int) []int {
	length := len(nums)

	forward := make([]int, length)
	backward := make([]int, length)
	res := make([]int, length)

	productForward := 1
	productBackward := 1

	for i := range nums {
		productForward *= nums[i]
		forward[i] = productForward

		productBackward *= nums[length-(i+1)]
		backward[length-(i+1)] = productBackward
	}

	for i := range nums {
		if i == 0 {
			res[i] = backward[i+1]
			continue
		}

		if i == length-1 {
			res[i] = forward[i-1]
			continue
		}

		res[i] = forward[i-1] * backward[i+1]

	}

	return res

}
