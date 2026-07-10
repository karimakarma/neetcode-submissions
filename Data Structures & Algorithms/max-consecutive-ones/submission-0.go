func findMaxConsecutiveOnes(nums []int) (maxCount int) {
	count := 0
	for i := range nums {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return
}
