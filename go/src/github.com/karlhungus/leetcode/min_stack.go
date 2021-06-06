package main

type MinStack struct {
	min int
	stk []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stk) == 0 || this.min > val {
		this.min = val
	}
	this.stk = append(this.stk, val)
}

func (this *MinStack) Pop() {
	val := this.Top()
	this.stk = this.stk[:len(this.stk)-1]
	if this.min == val && len(this.stk) > 0 {
		this.min = this.stk[0]
		for _, i := range this.stk {
			if this.min > i {
				this.min = i
			}
		}
	}

}

func (this *MinStack) Top() int {
	val := this.stk[len(this.stk)-1]
	return val
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
