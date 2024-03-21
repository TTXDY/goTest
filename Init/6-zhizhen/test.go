package main

func ch(p *int) {
	*p = 9
}

func main() {
	var a = 10
	var b *int
	//*b = 20
	ch(&a)
	b = &a
	//ch(b)
	println(a)
	println(&a)
	print(*b)
}
