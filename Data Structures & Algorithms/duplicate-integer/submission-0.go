
func hasDuplicate(nums []int) bool {
	slice := map[int]bool{}

	for _, x := range nums {
		if slice[x] {
			return true
		}
		slice[x] = true
	}

	return false
}
