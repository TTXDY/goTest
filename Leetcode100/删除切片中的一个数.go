package main

import "fmt"

func remove(num []int, i int) []int {
	copy(num[i:], num[i+1:])
	return num[:len(num)-1]
}

func main() {
	stack := make([]int, 10)
	for i := 0; i < 10; i++ {
		stack[i] = i
	}
	remove(stack, 2-1)
	fmt.Println(stack)
}
