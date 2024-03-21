package main

// iota 只能够配合const（）一起使用，iota只有在const进行累加效果。
const (
	BEIJING = iota * 10
	SHANGHai
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, u = iota * 10, iota * 20
	t, h
)

func main() {
	const a int = 100
	println("a = ", a)
	println("shanghai = ", SHANGHai)
}
