func maxProfit(prices []int) (res int) {
	for buy := 0; buy < len(prices)-1; buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			if prices[sell] > prices[buy] {
				res = max(res, prices[sell]-prices[buy])
			}
		}
	}

	return
}
