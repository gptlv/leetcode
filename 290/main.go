package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat cat"))
}

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}

	pm := make(map[byte]int)
	wm := make(map[string]int)

	for idx := range pattern {
		if pm[pattern[idx]] != wm[words[idx]] {
			return false
		}
		pm[pattern[idx]], wm[words[idx]] = idx+1, idx+1
	}

	return true
}
