func topKFrequent(nums []int, k int) (res []int) {
	seenNumbs := map[int]int{}

	for _, num := range nums {
		seenNumbs[num]++
	}

	rank := make([][]int, len(nums)+1)
	for n, f := range seenNumbs {
		rank[f] = append(rank[f], n)
	}

	for i := len(rank) - 1; i >= 0; i-- {
		for _, n := range rank[i] {
			res = append(res, n)

			if len(res) == k {
				return
			}
		}
	}

	return
}
