package main

func main() {

}

func lengthOfLastWord(s string) int {
	length := 0
	ptr := len(s) - 1
	for ptr >= 0 {
		if s[ptr] != ' ' {
			length++
		} else if length > 0 {
			break
		}
		ptr--
	}

	return length
}
