type Stack[T any] struct {
	data   []T
	length int
}

func (st *Stack[T]) pop() T {
	if st.length > 0 {
		st.length--
		val := st.data[st.length]
		st.data = st.data[:st.length]
		return val
	}

	var def T
	return def
}

func (st *Stack[T]) push(item T) {
	st.length++
	st.data = append(st.data, item)
}

func (st Stack[T]) peek() T {
	return st.data[st.length-1]
}

func largestRectangleArea(heights []int) (maxArea int) {
	heights = append(heights, 0)
	stack := Stack[int]{}

	for i, height := range heights {

		for stack.length != 0 && heights[stack.peek()] > height {
			h := heights[stack.pop()]
			w := i
			if stack.length > 0 {
				w = i - stack.peek() - 1
			}

			area := h * w
			if area > maxArea {
				maxArea = area
			}

		}
		stack.push(i)
	}

	return
}
