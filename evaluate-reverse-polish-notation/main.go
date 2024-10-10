package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	s := []int{}

	for i := 0; i < len(tokens); i++ {
		if isOperator(tokens[i]){
            if len(s) < 2 {
                return 0
            }

			b := s[len(s)-1]
			a := s[len(s)-2]
			s = s[:len(s)-2]

			val := calculate(a, b, tokens[i])
			s = append(s, val)
		} else {
            num, _ := strconv.Atoi(tokens[i])
            s = append(s, num)
        }
	}

	return s[0]
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return -1
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	}
	return false
}
