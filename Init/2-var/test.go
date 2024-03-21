package main

import "fmt"

func main() {
	var a int
	fmt.Println("a", a)
	var b int = 100
	fmt.Println("b", b)
	var c = 100
	fmt.Printf("type of c is %T", c)
	println("")
	fmt.Printf("c = %d, type of c is %T\n", c, c)

	// mod只能使用在函数体内
	e := 100
	fmt.Printf("e = %d, type of d is %T\n", e, e)

	f := "abchd"
	fmt.Printf("f = %s, type of f is %T\n", f, f)

	h := 3.14
	fmt.Println("h = ", h)
	fmt.Printf("type of f is %T\n", h)
}
