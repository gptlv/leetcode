package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

// func maxProfit(prices []int) int {
// 	ptr := 0
// 	profit := 0

// 	for delta := 0; ptr < len(prices); delta++ {
// 		if delta > len(prices)-ptr-1 {
// 			ptr++
// 			delta = 0
// 			continue
// 		}

// 		if prices[ptr+delta]-prices[ptr] > profit {
// 			profit = prices[ptr+delta] - prices[ptr]
// 		}
// 	}

//		return profit
//	}
func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}
