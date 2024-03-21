package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 232. 用栈实现队列 10min
* @ Author: Jay
* @ Date: 2024/1/17
* @ Time: 15:18
 */

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	queue := MyQueue{
		stack: make([]int, 0),
	}
	return queue
}

// Push 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

// Pop 从队列的开头移除并返回元素
func (this *MyQueue) Pop() int {
	x := this.stack[0]
	this.stack = this.stack[1:]
	return x
}

// Peek 返回队列开头的元素
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

// Empty 如果队列为空，返回 true ；否则，返回 false
func (this *MyQueue) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(10)
	m := obj.Pop()
	fmt.Println(m)
}
