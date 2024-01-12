package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}

func isIsomorphic(s string, t string) bool {
	sm := make(map[byte]int)
	tm := make(map[byte]int)

	for idx := range s {
		if sm[s[idx]] != tm[t[idx]] {
			return false
		}
		sm[s[idx]] = idx + 1
		tm[t[idx]] = idx + 1
	}

	return true
}
