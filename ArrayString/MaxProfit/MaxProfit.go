package maxprofit

import "fmt"

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		if currentPrice > buyPrice {
			profit += currentPrice - buyPrice
		} else {
			buyPrice = currentPrice
		}
		fmt.Println(buyPrice)
	}

	return profit
}
