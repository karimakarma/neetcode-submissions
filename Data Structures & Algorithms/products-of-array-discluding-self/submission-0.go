func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))

	for i := range nums {
		x := 1
		for idx, num := range nums {
			if i != idx {
				x *= num
			}
		}
		res = append(res, x)
	}

	return res
}
