func isValid(s string) bool {
	var stack []rune
    paran_map := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, paran := range s {
		if val, ok := paran_map[paran]; ok {
			stack = append(stack, val)
		} else if len(stack) > 0 && stack[len(stack) - 1] == paran {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
