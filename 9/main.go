package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(10000021))
}

func isPalindrome(x int) bool {

	// fmt.Println(x)

	if x < 0 {
		return false
	}

	len := getNumberLength(x)

	if len == 1 {
		return true
	}

	if len == 2 {
		return x/10 == x%10
	}

	firstDigit := x / powInt(10, len-1)
	lastDigit := x % 10

	// fmt.Println(firstDigit, lastDigit)

	if firstDigit == lastDigit {
		digitsInTheMiddle := (x - (firstDigit * powInt(10, len-1))) / 10
		fmt.Println(digitsInTheMiddle, x)
		return isPalindrome(digitsInTheMiddle)
	}

	return false
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func getNumberLength(x int) int {
	// len := 1

	// for q := x; q > 9; {
	// 	q = q / 10
	// 	len++
	// }
	// return len

	return 1 + int(math.Log10(float64(x)))

}
