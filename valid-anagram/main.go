package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[rune]int)
	tm := make(map[rune]int)

	for _, ch := range s {
		sm[ch]++
	}

	for _, ch := range t {
		tm[ch]++
	}

	for k := range sm {
		if _, ok := tm[k]; !ok {
			return false
		}

		if sm[k] != tm[k] {
			return false
		}
	}

	return true
}
