func evalRPN(tokens []string) int {
	var stack []int
	for _, tok := range tokens {
		switch tok {
			case "*", "/", "+", "-":
			a, b := stack[len(stack) - 2], stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]

			switch tok {
				case "*":
				stack = append(stack, a * b)
				case "/":
				stack = append(stack, a / b)
				case "+":
				stack = append(stack, a + b)
				case "-":
				stack = append(stack, a - b)
			}

			default:
			num, _ := strconv.Atoi(tok)
			stack = append(stack, num)
		}
	}

	return stack[len(stack) - 1]
}
