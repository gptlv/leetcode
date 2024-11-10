package main

func main() {

}

func plusOne(digits []int) []int {
	idx := len(digits) - 1

	for idx >= 0 {
		if digits[idx] < 9 {
			digits[idx]++
			return digits
		}
		digits[idx] = 0
		idx--

	}

	return append([]int{1}, digits...)
}
