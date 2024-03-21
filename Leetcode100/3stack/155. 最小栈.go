package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 155. 最小栈
* @ Author: Jay
* @ Date: 2024/1/17
* @ Time: 15:36
 */

/*
实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	stack := MinStack{
		stack: make([]int, 0),
	}
	return stack
}

// Push void push(int val) 将元素val推入堆栈。
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

// Pop void pop() 删除堆栈顶部的元素。
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

// Top int top() 获取堆栈顶部的元素。
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin int getMin() 获取堆栈中的最小元素。
func (this *MinStack) GetMin() int {
	themin := this.stack[0]
	for _, val := range this.stack {
		if val < themin {
			themin = val
		}
	}
	return themin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
