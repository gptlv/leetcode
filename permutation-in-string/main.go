package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	freq1 := [26]int{}
	freq2 := [26]int{}
	left := 0

	for i := range s1 {
		freq1[s1[i]-'a']++
	}

	for right := 0; right < len(s2); right++ {
		freq2[s2[right]-'a']++

		if right-left+1 > len(s1) {
			freq2[s2[left]-'a']--
			left++
		}

		if freq1 == freq2 {
			return true
		}
	}

	return false
}
