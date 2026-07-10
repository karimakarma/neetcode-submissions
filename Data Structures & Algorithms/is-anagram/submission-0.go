
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	slice := [26]int{}

	for i := range s {
		slice[int(s[i]-97)]++
		slice[int(t[i]-97)]--
	}

	for _, i := range slice {
		if i != 0 {
			return false
		}
	}
	return true
}
