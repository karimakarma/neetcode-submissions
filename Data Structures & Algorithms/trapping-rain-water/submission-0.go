func trap(height []int) (res int) {
	var maxL, maxR int
	for l, r := 0, len(height)-1; l < r; {
		if height[r] < height[l] {
			maxR = max(maxR, height[r])
			res += maxR - height[r]
			r--
		} else {
			maxL = max(maxL, height[l])
			res += maxL - height[l]
			l++
		}
	}

	return
}
