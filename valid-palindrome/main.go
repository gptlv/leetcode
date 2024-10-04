package main

import (
	"strings"
)

func main() {

}

func isPalindrome(s string) bool {
	low := strings.ToLower(s)
	str := []byte{}
	for i := 0; i < len(low); i++ {
		if low[i] <= 122 && low[i] >= 97 || low[i] <= 57 && low[i] >= 48 {
			str = append(str, low[i])
		}

		if low[i] <= 90 && low[i] >= 65 {
			str = append(str, low[i]+32)
		}
	}

	left, right := 0, len(str)-1

	for left <= len(str)/2 && right >= len(str)/2 && len(str) != 0 {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}
