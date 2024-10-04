package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, ch := range magazine {
		m[ch]++
	}

	for _, ch := range ransomNote {
		val, ok := m[ch]
		if !ok || val == 0 {
			return false
		}

		m[ch]--
	}

	return true
}
