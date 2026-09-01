type MinStack struct {
	min int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min: 0,
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
		this.stack = append(this.stack, 0)
	} else {
		this.stack = append(this.stack, val - this.min)

		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack) - 1]
	if top < 0 {
		this.min = this.min - top
	}

	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack) - 1]
	if top > 0 {
		return this.min + top
	}
	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
