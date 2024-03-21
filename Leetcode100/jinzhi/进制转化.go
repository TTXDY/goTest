package main

import "fmt"

func twoToTen(n int) int {
	list := make(map[int]int)
	if n == 0 {
		list[0] = 0
	}
	if n == 1 {
		list[0] = 1
	}
	if n == 2 {
		list[0] = 0
		list[1] = 1
	}
	for i := 0; n != 0; i++ {
		list[i] = n % 2
		n = n / 2
	}
	var sum int = 0
	for _, v := range list {
		sum = sum + v
	}
	//fmt.Println(1list)
	return sum

}

func countBits(n int) []int {
	list := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		list[i] = twoToTen(i)
	}
	return list
}

func main() {
	a := 5
	fmt.Println(countBits(a))
}
