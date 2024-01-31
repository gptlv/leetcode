package main

func main() {

}

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	var delta int
	var profit int

	for i := 0; i < len(prices)-1; i++ {
		delta = prices[i+1] - prices[i]

		if delta > 0 {
			profit += delta
		}
	}

	return profit

}
