type MinStack struct {
	main_s, min_s []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.main_s = append(this.main_s, val)

	if len(this.min_s) == 0 || val <= this.min_s[len(this.min_s) - 1] { this.min_s = append(this.min_s, val) }
}

func (this *MinStack) Pop() {
	if this.GetMin() == this.main_s[len(this.main_s) - 1] { this.min_s = this.min_s[:len(this.min_s) - 1] }

	this.main_s = this.main_s[:len(this.main_s) - 1]
}

func (this *MinStack) Top() int {
	return this.main_s[len(this.main_s) - 1]
}

func (this *MinStack) GetMin() int {
	return this.min_s[len(this.min_s) - 1]
}

