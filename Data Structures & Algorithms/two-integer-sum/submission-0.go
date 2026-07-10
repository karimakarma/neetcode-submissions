func twoSum(nums []int, target int) []int {
	seenMap := map[int]int{}

	for idx, num := range nums {
		x := target - num

		if val, ok := seenMap[x]; ok {
			return []int{val, idx}
		}
		seenMap[num] = idx

	}

	return []int{}
}
