type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.mins[len(this.mins)-1] == this.stack[len(this.stack)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
