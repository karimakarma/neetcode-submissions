func longestConsecutive(nums []int) int {
	seenNums := map[int]bool{}

	for _, num := range nums {
		seenNums[num] = true
	}

	maxCount := 0
	for _, num := range nums {
		if !seenNums[num-1] {
			numCopy := num
			count := 1

			for seenNums[numCopy+1] {
				numCopy++
				count++
			}
			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}
