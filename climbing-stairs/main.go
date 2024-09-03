package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	sequence := getSequence(n)
	return sequence[n-1] + sequence[n-2]

}

func getSequence(n int) []int {
	sequence := make([]int, n+1)

	sequence[0] = 1
	sequence[1] = 1

	for i := 2; i < n; i += 1 {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}
