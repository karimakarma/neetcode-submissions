func removeElement(nums []int, val int) int {
	pointer := 0
	for i := range nums {
		if nums[i] != val {
			nums[pointer] = nums[i]
			pointer++
		}
	}

	return pointer
}