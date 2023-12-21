package main

func main() {

}

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0, 10),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for _, x := range this.stack {
		if min > x {
			min = x
		}
	}

	return min
}
