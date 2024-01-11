package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	r := strings.ToLower(re.ReplaceAllString(s, ""))
	if r == "" {
		return true
	}

	lp := 0
	rp := len(r) - 1

	for lp < rp {
		if r[lp] != r[rp] {
			return false
		}
		lp++
		rp--
	}

	return true
}
