package leetcode

func MaxProfit(prices []int) int {
	buy := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < buy {
			buy = price
		} else {
			profit := price - buy
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
