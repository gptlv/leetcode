package main

func main() {

}

func moveZeroes(nums []int) {
	var ptr, zeros int

	for _, num := range nums {
		if num != 0 {
			nums[ptr] = num
			ptr++
		} else {
			zeros++
		}
	}

	for i := len(nums) - zeros; i < len(nums); i++ {
		if i < 0 {
			continue
		}
		nums[i] = 0
	}

}
