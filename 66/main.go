package main

func main() {

}

func plusOne(digits []int) []int {
	idx := len(digits) - 1
	shouldAppend := false
	digits[idx]++

	for idx >= 0 {
		if shouldAppend {
			digits[idx]++
			shouldAppend = false
		}
		if digits[idx] == 10 {
			digits[idx] = 0
			shouldAppend = true
		}
		idx--

	}

	if shouldAppend {
		return append([]int{1}, digits...)
	}

	return digits
}
