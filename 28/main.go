package main

import "fmt"

func main() {

	fmt.Println(strStr("abc", "c"))

}

func strStr(haystack string, needle string) int {
	if haystack == "" || needle == "" {
		return -1
	}
	nl := len(needle)
	hl := len(haystack)

	if nl > hl {
		return -1
	}

	for i := 0; i < hl-nl+1; i++ {
		hss := haystack[i : i+nl]
		fmt.Println(hss)
		if hss == needle {
			return i
		}
	}
	return -1
}
