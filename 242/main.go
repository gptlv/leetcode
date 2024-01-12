package main

import "fmt"

func main() {

	fmt.Println(isAnagram("rat", "cat"))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := make(map[rune]int)
	tm := make(map[rune]int)

	for _, sr := range s {
		sm[sr]++
	}
	for _, tr := range t {
		tm[tr]++
	}

	for key, value := range tm {
		if _, ok := sm[key]; !ok {
			return false
		}
		if sm[key] != value {
			return false
		}
	}

	return true

}
