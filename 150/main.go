package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	noteMap := make(map[rune]int)

	for _, char := range magazine {
		magazineMap[char]++
	}

	for _, char := range ransomNote {
		noteMap[char]++
		if _, ok := magazineMap[char]; !ok {
			return false
		}
		if magazineMap[char] < noteMap[char] {
			return false
		}
	}

	return true
}
