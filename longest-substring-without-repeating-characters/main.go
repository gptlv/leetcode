package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]bool)
	left := 0
	res := 0

	for right := 0; right < len(s); right++ {
		for m[s[right]] {
			delete(m, s[left])
			left++
		}

		m[s[right]] = true
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

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 || len(s) == 1 {
// 		return len(s)
// 	}

// 	left := 0
// 	res := 0

// 	for i := 0; i < len(s); i++ {
// 		right := i + 1
// 		unique := countUnique(s[left:right])

// 		if unique < right-left {
// 			left++
// 		} else {
// 			res = max(res, unique)
// 		}
// 	}

// 	return res
// }

// func countUnique(window string) int {
// 	m := make(map[uint8]struct{})
// 	for i := range window {
// 		m[window[i]] = struct{}{}
// 	}

// 	return len(m)
// }
