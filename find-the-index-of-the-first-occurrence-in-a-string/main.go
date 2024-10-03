package main

func main() {

}

func strStr(haystack string, needle string) int {
	start, end := 0, len(needle)

	for end <= len(haystack) {
		if haystack[start:end] == needle {
			return start
		}

		start++
		end++
	}

	return -1
}
