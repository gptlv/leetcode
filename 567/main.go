package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	freq := getFreq(s1)

	w := len(s1)

	lp := 0
	rp := lp + w

	for rp <= len(s2) {
		if freq == getFreq(s2[lp:rp]) {
			return true
		}

		lp++
		rp = lp + w
	}

	return false
}

func getFreq(s string) [26]int {
	freq := [26]int{}
	for _, r := range s {
		freq[r-'a']++
	}
	return freq
}
