func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	for i := range temperatures {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}
