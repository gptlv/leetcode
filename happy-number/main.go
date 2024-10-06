package main

import (
	"strconv"
	"strings"
)

func main() {

}

func isHappy(n int) bool {
	m := make(map[int]int)
	sum := 0

	for {
		nStr := strconv.Itoa(n)
		digits := strings.Split(nStr, "")
		sum = 0

		for _, digit := range digits {
			d, _ := strconv.Atoi(digit)
			sum += d * d
		}

		if sum == 1 {
			return true
		}

		if _, ok := m[sum]; ok {
			return false
		}

		m[n] = sum
		n = sum
	}
}
