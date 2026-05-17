package besttimetobuyandsellstock

// MaxProfit calculates max profit from stock prices.
// Time O(n), since we should iterate all the input
// Space O(1), sine we don't allocate any additional memory.
func MaxProfit(prices []int) int {
	profit, currPrice := 0, prices[0]

	for i := range prices {
		currPrice = min(currPrice, prices[i])
		profit = max(profit, prices[i]-currPrice)
	}

	return profit
}
