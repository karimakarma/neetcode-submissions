func twoSum(numbers []int, target int) []int {
	seenMap := map[int]int{}

	for idx, num := range numbers {
		x := target - num

		if val, ok := seenMap[x]; ok {
			return []int{val, idx + 1}
		}
		seenMap[num] = idx + 1

	}

	return []int{}

}
