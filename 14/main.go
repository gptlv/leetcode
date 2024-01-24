package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, str := range strs {
		if len(str) < len(prefix) {
			prefix = str
		}
	}

	idx := 0

	for prefix != "" {
		if strs[idx][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			idx = 0
			continue
		}
		if idx == len(strs)-1 {
			break
		}
		idx++
	}

	return prefix
}
