func calPoints(operations []string) int {
	op := []int{}
	opMap := map[string]func(){
		"+": func() {
			op = append(op, op[len(op)-2]+op[len(op)-1])
		},
		"D": func() {
			op = append(op, 2*op[len(op)-1])
		},
		"C": func() {
			op = op[:len(op)-1]
		},
	}

	for _, s := range operations {
		if fn, ok := opMap[s]; ok {
			fn()
		} else {
			x, _ := strconv.Atoi(s)
			op = append(op, x)
		}
	}

	sum := 0
	for _, n := range op {
		sum += n
	}

	return sum
}