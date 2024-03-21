package main

import "fmt"

func main() {
	stack := make([]int, 10)
	for i := 0; i < 10; i++ {
		stack[i] = i
	}
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
