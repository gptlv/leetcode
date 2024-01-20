package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		arr := [26]int{}

		for _, r := range str {
			arr[r-'a']++
		}

		m[arr] = append(m[arr], str)
	}

	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}

	return res

	// initial solution

	// m := make(map[string][]string)

	// for _, str := range strs {
	// 	chars := []rune(str)

	// 	sort.Slice(chars, func(i, j int) bool {
	// 		return chars[i] < chars[j]
	// 	})

	// 	m[string(chars)] = append(m[string(chars)], str)

	// }

	// res := make([][]string, 0, len(m))

	// for _, v := range m {
	// 	res = append(res, v)
	// }

	// return res
}
