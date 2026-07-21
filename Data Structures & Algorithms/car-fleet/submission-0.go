type Car struct {
	position int
	time float32
}
func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i := range position {
		cars[i] = Car{position[i], float32(target-position[i]) / float32(speed[i])}
	}

	sort.Slice(cars, func(i int, j int) bool { return cars[i].position > cars[j].position })

	var stack []float32
	for _, c := range cars {
		stack = append(stack, c.time)

		if len(stack) >= 2 {
			if stack[len(stack)-1] <= stack[len(stack)-2] {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack)
}
