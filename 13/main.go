package main

import "fmt"

func main() {

	str := "MCMXCIV"
	fmt.Println(romanToInt(str))
	// fmt.Println(isSubstractable('I'))

}

func isSubtractable(r rune) bool {
	for _, numeral := range "IXC" {
		if numeral == r {
			return true
		}
	}
	return false
}

func getSubtractedNumber(n string) int {

	fmt.Println(n)

	subtractable := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	number, ok := subtractable[n]
	if !ok {
		return 0
	}

	return number
}

func romanToInt(s string) int {

	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int

	for i := 0; i < len(s); i++ {
		if isSubtractable(rune(s[i])) && i != len(s)-1 {
			number := getSubtractedNumber(string(s[i]) + string(s[i+1]))
			if number != 0 {
				sum += number
				i++
				continue
			}
		}
		sum += romanNumerals[rune(s[i])]
	}
	return sum
}
