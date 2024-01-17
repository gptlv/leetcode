package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	n := x
	r := 0

	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}

	return r == x
}
