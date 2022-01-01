package main

import (
	"math"
)

/*
	push(x)  —— 将元素 x 推入栈中。
	pop()    —— 删除栈顶的元素。
	top()    —— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。
*/
type MinStack struct {
	data   []int
	minNum int
}

func Constructor() MinStack {
	return MinStack{
		data:   make([]int, 0, 1024),
		minNum: math.MaxInt64,
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if this.minNum != math.MaxInt64 && val < this.minNum {
		this.minNum = val
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		topNum := this.data[len(this.data)-1]
		if topNum == this.minNum {
			this.minNum = math.MaxInt64
		}
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.minNum != math.MaxInt64 {
		return this.minNum
	}
	for _, num := range this.data {
		if num < this.minNum {
			this.minNum = num
		}
	}
	return this.minNum
}

func main() {
	//  ["MinStack","push","push","getMin","getMin","push","getMin","getMin","top","getMin","pop","push","push","getMin","push","pop","top","getMin","pop"]
	//	[[],[-10],[14],[],[],[-20],[],[],[],[],[],[10],[-7],[],[-7],[],[],[],[]]
	stack := Constructor()
	stack.Push(-10)
	stack.Push(-14)
	stack.GetMin()
	stack.GetMin()
	stack.Push(-20)
	stack.GetMin()
	stack.GetMin()
	stack.Top()
	stack.GetMin()
	stack.Pop()
	stack.Push(10)
	stack.Push(-7)
	stack.GetMin()
	stack.Push(-7)
	stack.Pop()
	stack.Top()
	stack.GetMin()
	stack.Pop()
}
