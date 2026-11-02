func evalRPN(tokens []string) int {
	var stack []int
	for _, tok := range tokens {
		if num, e := strconv.Atoi(tok); e != nil { 
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
		} else {
			stack = append(stack, num) 
		} 
	}

	return stack[0]
}
