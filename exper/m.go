package main

import (
	"fmt"
	"math"
)

func Print(m []int) {
	//for i := 0; i < 10; i++ {
	//	println(m[i])
	//}
	for k, v := range m {
		println("k is ", k, "v is ", v)
	}
}

func main() {
	//3stack := make([]int, 2)
	//Print(3stack)
	//fmt.Printf("%T\n", 3stack)
	//3stack = append(3stack, 10)
	//Print(3stack)
	fmt.Println(math.MaxInt32)
}
