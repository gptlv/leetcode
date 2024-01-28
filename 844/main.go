package main

import "fmt"

func main() {

}

func backspaceCompare(s string, t string) bool {

	s = handleBackspace(s)
	t = handleBackspace(t)

	fmt.Println(s, t)

	return s == t

}

func handleBackspace(s string) string {

	res := make([]byte, len(s))
	lp := 0
	rp := 0

	for rp < len(s) {
		if s[rp] == '#' {
			lp--
		} else {
			res[lp] = s[rp]
			lp++
		}
		rp++

		if lp < 0 {
			lp = 0
		}
	}

	return string(res[:lp])
}
