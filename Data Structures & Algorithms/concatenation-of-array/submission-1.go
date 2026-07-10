func getConcatenation(nums []int) []int {
	length := len(nums)
	res := make([]int, length*2)

	for i := range nums {
		res[i] = nums[i]
		res[i+length] = nums[i]
	}

	return res

}
