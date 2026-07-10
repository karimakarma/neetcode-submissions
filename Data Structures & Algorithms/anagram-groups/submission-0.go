import "slices"
func groupAnagrams(strs []string) [][]string {
	sortStr := func(s string) string {
		slice := strings.Split(s, "")
		slices.Sort(slice)

		s = strings.Join(slice, "")

		return s
	}

	seenWord := map[string][]int{}

	for idx, str := range strs {
		str = sortStr(str)

		seenWord[str] = append(seenWord[str], idx)
	}

	res := make([][]string, len(seenWord))
	
	i := 0
	for _, s := range seenWord {
		for _, idx := range s {
			res[i] = append(res[i], strs[idx])
		}
		i++
	}

	return res
}
