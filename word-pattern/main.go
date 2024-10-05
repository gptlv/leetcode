package main

import "strings"

func main() {

}

func wordPattern(pattern string, s string) bool {
	m1 := make(map[uint8]string)
	m2 := make(map[string]uint8)

	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		m1[pattern[i]] = words[i]
	}

	for i := 0; i < len(pattern); i++ {
		if m1[pattern[i]] != words[i] {
			return false
		}
	}

	for i := 0; i < len(words); i++ {
		m2[words[i]] = pattern[i]
	}

	for i := 0; i < len(words); i++ {
		if m2[words[i]] != pattern[i] {
			return false
		}
	}

	return true
}
