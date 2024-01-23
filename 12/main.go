package main

import "strings"

func main() {

}

func intToRoman(num int) string {

	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	k := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	idx := 0

	var sb strings.Builder

	for num != 0 {
		if num/k[idx] == 0 {
			idx++
			continue
		}

		sb.WriteString(s[idx])
		num -= k[idx]
	}

	return sb.String()
}
