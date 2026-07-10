func maxProfit(prices []int) (res int) {
	buy := prices[0]

	for _, price := range prices {
		if price < buy {
			buy = price
		} else if price-buy > res {
			res = price - buy
		}
	}

	return
}

