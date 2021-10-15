package Q121_Best_Time_to_Buy_and_Sell_Stock

//Runtime: 116 ms, faster than 95.81% of Go online submissions for Best Time to Buy and Sell Stock.
//Memory Usage: 8.3 MB, less than 49.28% of Go online submissions for Best Time to Buy and Sell Stock.
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 1 {
		return 0
	}
	minPrice, max := prices[0], 0
	for i := 1; i < l; i++ {
		if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return max
}
