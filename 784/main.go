package main

func main() {

}

func letterCasePermutation(s string) []string {
	idx := 0
	res := []string{}
	cur := ""

	helper(s, idx, cur, &res)

	return res
}

func helper(s string, idx int, cur string, res *[]string) {
	if idx == len(s) {
		*res = append(*res, cur)
		return
	}

	if s[idx] >= 'A' && s[idx] <= 'Z' {
		helper(s, idx+1, cur+string(s[idx]+32), res)
	} else if s[idx] >= 'a' && s[idx] <= 'z' {
		helper(s, idx+1, cur+string(s[idx]-32), res)
	}

	helper(s, idx+1, cur+string(s[idx]), res)
}
