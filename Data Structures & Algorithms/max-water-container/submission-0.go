func maxArea(heights []int) (maxA int) {
	for l, r := 0, len(heights)-1; l < r; {
		h := min(heights[l], heights[r])

		maxA = max(h*(r-l), maxA)

		if h == heights[l] {
			l++
		} else {
			r--
		}
	}

	return
}
