package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		freq := getFrequency(str)
		m[freq] = append(m[freq], str)
	}

	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func getFrequency(str string) [26]int {
	freq := [26]int{}
	for _, ch := range str {
		freq[ch-'a']++
	}

	return freq
}
