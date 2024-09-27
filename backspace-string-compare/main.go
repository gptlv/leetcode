package main

func main() {

}

func backspaceCompare(s string, t string) bool {
	return helper(s) == helper(t)
}

func helper(s string) string {
	res := make([]byte, len(s))
	ptr1, ptr2 := 0, 0

	for ptr2 < len(s) {
		if s[ptr2] != '#' {
			res[ptr1] = s[ptr2]
			ptr1++
		} else {
			ptr1--
		}

		ptr2++

		if ptr1 < 0 {
			ptr1 = 0
		}
	}

	return string(res[:ptr1])
}
