package main

func main() {

}

func letterCombinations(digits string) []string {
	m := make(map[byte]string, 8)
	res := []string{}
	if len(digits) == 0 {
		return res
	}

	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"

	var backtrack func(i int, combination []byte)

	backtrack = func(i int, combination []byte) {
		if len(combination) == len(digits) {
			c := make([]byte, len(digits))
			copy(c, combination)
			res = append(res, string(c))
			return
		}

		if i >= len(digits) {
			return
		}

		letters := m[digits[i]]
		for idx := range letters {
			letter := letters[idx]
			combination = append(combination, letter)
			backtrack(i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0, []byte{})

	return res
}
