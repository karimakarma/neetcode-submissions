func evalRPN(tokens []string) int {
	nums := []int{}
	for i := range tokens {
		switch tokens[i] {
		case "+":
			x := nums[len(nums)-2] + nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, x)
		case "-":
			x := nums[len(nums)-2] - nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, x)
		case "*":
			x := nums[len(nums)-2] * nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, x)
		case "/":
			x := nums[len(nums)-2] / nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, x)
		default:
			n, _ := strconv.Atoi(tokens[i])
			nums = append(nums, n)
		}
	}
	return nums[0]
}
