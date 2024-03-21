package main

import "fmt"

func Print(a int, c int, b string) int {
	d := a + c
	println("I come from: ", b)
	return d
}

func Print2(a int, c int, b string) (int, int) {
	d := a + c
	x := a - c
	println("I come from: ", b)
	return d, x
}

func Print3(a int, c int, b string) (d int, x int) {
	d = a + c
	x = a - c
	println("I come from: ", b)
	return
}
func main() {
	a := "China"
	println("10+20 =", Print(10, 20, a))
	m, n := Print2(10, 20, a)
	println(n)
	fmt.Printf("type of Print2 is %T", m)
	m1, n1 := Print3(10, 20, a)
	println(n1)
	fmt.Printf("type of Print2 is %T", m1)
}
