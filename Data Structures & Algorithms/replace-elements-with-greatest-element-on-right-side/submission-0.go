func replaceElements(arr []int) []int {
	maxR := -1
	for i := len(arr) - 1; i >= 0; i-- {
		p := arr[i]
		arr[i] = maxR
		if p > maxR {
			maxR = p
		}
	}

	return arr
}
