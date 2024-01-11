package main

import "fmt"

func main() {
	fmt.Println(isValid("(){}[]"))

}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	idx := 1
	for {
		if s == "" {
			return true
		}
		if len(s) < 2 {
			return false
		}
		currentRune := rune(s[idx])
		previousRune := rune(s[idx-1])

		_, ok := brackets[currentRune]
		if !ok && idx == len(s)-1 {
			return false
		}
		if previousRune == brackets[currentRune] {
			s = s[:idx-1] + s[idx+1:]
			idx = 1
			continue
		}
		if !ok {
			idx++
			continue
		}
		return false
	}
}
