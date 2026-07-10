func lengthOfLongestSubstring(s string) (length int) {
	seenByte := map[byte]int{}
	l := 0
	for r := l; r < len(s); r++ {
		if val, ok := seenByte[s[r]]; ok {
			if val > l {
				l = val
			}
		}
		seenByte[s[r]] = r + 1
		if x := r - l + 1; x > length {
			length = x
		}
	}

	return
}
