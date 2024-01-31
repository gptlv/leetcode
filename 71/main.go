package main

import (
	"strings"
)

func main() {

	simplifyPath("/home/user/../../usr/local/bin")

}

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	res := []string{}

	for _, str := range s {
		if str != "" && str != "." {
			res = append(res, str)
		}
	}

	idx := 0

	for _, name := range res {
		if name == ".." {
			idx--
			if idx < 0 {
				idx = 0
			}
			continue
		}

		res[idx] = name
		idx++
	}

	return ("/" + strings.Join(res[:idx], "/"))

}
