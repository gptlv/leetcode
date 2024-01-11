package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord(" fdfd fdfdfdfdf fdfdf  fdfd          "))
}

func lengthOfLastWord(s string) int {
	s = (strings.TrimRight(s, " "))
	var l int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return l
		}
		l++
	}
	return l
}
