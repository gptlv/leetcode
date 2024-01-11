package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	profit := 0

	for _, price := range prices {
		if price < min {
			min = price
		}
		if d := price - min; d > profit {
			profit = d
		}
	}

	return profit
}
