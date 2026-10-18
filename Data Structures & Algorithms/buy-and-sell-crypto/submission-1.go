func maxProfit(prices []int) int {
	bestPrice := 0

	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices);j++{
			diff := prices[j]-prices[i]
			if diff>bestPrice{
				bestPrice = diff
			}
		}
	}
	return bestPrice
}
