package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	res := []string{}
	s = strings.TrimSpace(s)

	left, right := len(s)-1, len(s)-1

	for left >= 0 {
		if s[left] != ' ' && right == -1 {
			right = left
		}

		if s[left] == ' ' && right != -1 {
			res = append(res, s[left+1:right+1])
			right = -1
		}
		if left == 0 {
			res = append(res, s[:right+1])
		}

		left--
	}

	return strings.Join(res, " ")
}
