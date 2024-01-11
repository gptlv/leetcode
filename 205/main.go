package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[rune]rune)
	tm := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		sr := rune(s[i])
		tr := rune(t[i])
		if _, ok := sm[sr]; ok && sm[sr] != tr {
			return false
		}
		sm[rune(s[i])] = rune(t[i])
		if _, ok := tm[tr]; ok && tm[tr] != sr {
			return false
		}
		tm[rune(t[i])] = rune(s[i])
	}

	for key, value := range sm {
		fmt.Println(string(key), string(value))
	}
	fmt.Println("///")
	for key, value := range tm {
		fmt.Println(string(key), string(value))
	}

	return len(sm) == len(tm)

}
