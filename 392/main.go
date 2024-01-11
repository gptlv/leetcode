package main

import "fmt"

func main() {

	fmt.Println(isSubsequence("b", "abc"))

}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}
	idx := 0

	for _, c := range t {
		if idx > len(s)-1 {
			break
		}
		if c == rune(s[idx]) {
			fmt.Println(c)
			idx++
		}
	}

	if idx == len(s) {
		return true
	}

	return false

}
