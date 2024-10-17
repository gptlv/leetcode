package main

func main() {

}

func characterReplacement(s string, k int) int {
	left := 0
	res := 0
	freqMap := make(map[uint8]int)
	maxFreq := 0

	for right := 0; right < len(s); right++ {
		freqMap[s[right]]++
		maxFreq = max(maxFreq, freqMap[s[right]])

		if (right-left+1)-maxFreq > k {
			freqMap[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
