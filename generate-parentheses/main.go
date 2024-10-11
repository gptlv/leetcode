package main

func main() {

}

func generateParenthesis(n int) []string {
	res := []string{}
	helper([]byte{}, n, n, &res)

	return res
}

func helper(q []byte, open, closed int, res *[]string) {
	if open == 0 && closed == 0 {
		*res = append(*res, string(q))
	}

	if open > 0 {
		helper(append(q, '('), open-1, closed, res)
	}

	if closed > open {
		helper(append(q, ')'), open, closed-1, res)
	}
}
