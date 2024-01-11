package main

import "fmt"

func main() {

	fmt.Println(climbStairs(1))
}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n < 3 {
		return n
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n-1]
}
