package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:225. 用队列实现栈 7min
* @ Author: Jay
* @ Date: 2024/1/17
* @ Time: 15:28
 */

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	stack := MyStack{stack: make([]int, 0)}
	return stack
}

// Push 将元素 x 压入栈顶。
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

// Pop 移除并返回栈顶元素。
func (this *MyStack) Pop() int {
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return x
}

// Top 返回栈顶元素。
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// Empty  如果栈是空的，返回 true ；否则，返回 false 。
func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {

}
