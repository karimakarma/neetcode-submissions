func maxProfit(prices []int) (res int) {
	buy := prices[0]

	for _, price := range prices {
		buy = min(price, buy)
		res = max(res, price-buy)
	}

	return
}
