package main

func main() {

}

func partition(s string) [][]string {
	res := [][]string{}
	var dfs func(int, []string)

	dfs = func(start int, part []string) {
		if start == len(s) {
			c := make([]string, len(part))
			copy(c, part)
			res = append(res, c)
			return
		}

		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				part = append(part, s[start:end])
				dfs(end, part)
				part = part[:len(part)-1]
			}
		}
	}

	dfs(0, []string{})
	return res
}

func isPalindrome(sub string) bool {
	left, right := 0, len(sub)-1
	for left < right {
		if sub[left] != sub[right] {
			return false
		}
		left++
		right--
	}
	return true
}
