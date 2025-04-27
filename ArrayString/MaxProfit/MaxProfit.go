package maxprofit

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		if currentPrice > buyPrice {
			currentProfit := currentPrice - buyPrice
			if currentProfit > profit {
				profit = currentProfit
			}
		} else {
			buyPrice = currentPrice
		}

	}

	return profit
}
