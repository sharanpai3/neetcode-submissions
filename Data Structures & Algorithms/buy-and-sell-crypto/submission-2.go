func maxProfit(prices []int) int {
	var l int
	r := l+1
	var bestPrice int

	for r< len(prices){
		if prices[r] > prices[l]{
		if prices[r]-prices[l] > bestPrice{
			bestPrice = prices[r]-prices[l]
		}} else{
			l = r
		}
			r++
	}
	return bestPrice
}
