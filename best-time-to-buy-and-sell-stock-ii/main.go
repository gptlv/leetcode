package main

func main() {

}

func maxProfit(prices []int) int {
	var profit int

	if len(prices) == 1 {
		return profit
	}

	start := 0

	for i := start + 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
