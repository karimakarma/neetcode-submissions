import "slices"
func threeSum(nums []int) (res [][]int) {
	slices.Sort(nums)

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, len(nums)-1; l < r; {
			if r == i {
				r--
			}

			sum := nums[l] + nums[r] + nums[i]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[l], nums[i], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}
