func findMaxConsecutiveOnes(nums []int) (maxCount int) {
	count := 0
	for i := range nums {
		if nums[i] == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return
}
