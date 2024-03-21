package main

import "fmt"

func twoToTen(n int) map[int]int {
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
	fmt.Println(list)
	return list

}
func hammingDistance(x int, y int) int {
	map1 := twoToTen(x)
	map2 := twoToTen(y)

	var sum int = 0
	for i := 0; i < max(len(map1), len(map2)); i++ {
		if map1[i] != map2[i] {
			sum++
		}
	}
	return sum
}

func main() {
	hammingDistance(1, 3)
}
