func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	template := [26]int{}
	window := [26]int{}

	for i := range s1 {
		template[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	if window == template {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		window[s2[i-len(s1)]-'a']--
		window[s2[i]-'a']++

		if window == template {
			return true
		}
	}

	return false
}
