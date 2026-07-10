func isPalindrome(s string) bool {
	isAlphanumeric := func(x byte) bool {
		return (x >= 'a' && x <= 'z') || (x >= '0' && x <= '9')
	}
	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}

		if !isAlphanumeric(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
