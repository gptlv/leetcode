package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	words := strings.Fields(s)

	res := []string{}

	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i])
	}

	return strings.Join(res, " ")
}
