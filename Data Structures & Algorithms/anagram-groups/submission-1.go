func groupAnagrams(strs []string) [][]string {
	seenWord := map[[26]byte][]string{}

	for _, str := range strs {

		var count [26]byte
		for i := range str {
			count[str[i]-'a']--
		}

		seenWord[count] = append(seenWord[count], str)
	}

	res := make([][]string, 0, len(seenWord))

	for _, slice := range seenWord {
		res = append(res, slice)
	}

	return res
}
