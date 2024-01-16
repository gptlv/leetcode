package main

import (
	"math"
)

func main() {

	countBits(5)

}

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	res := make([]int, n+1)

	res[0], res[1] = 0, 1

	for i := 2; i <= n; i++ {
		pow := math.Trunc(math.Log2(float64(i)))
		idx := i - int(math.Pow(float64(2), pow))

		res[i] = res[idx] + 1
	}

	return res
}
