func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	m := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	stack := make([]byte, 0, len(s)/2)

	for i := range s {
		if val, ok := m[s[i]]; ok {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
