package main

func main() {

}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	var sPtr, tPtr, count int
	for tPtr < len(t) && sPtr < len(s) {
		if t[tPtr] == s[sPtr] {
			count++
			sPtr++
		}

		tPtr++
	}

	return len(s) == count
}
