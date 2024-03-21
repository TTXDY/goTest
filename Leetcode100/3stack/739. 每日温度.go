package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 739. 每日温度
* @ Author: Jay
* @ Date: 2024/1/17
* @ Time: 15:59
 */

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var ansBiao int = 0
	for len(temperatures) != 0 {
		count := 0
		for i := 1; i < len(temperatures); i++ {
			if temperatures[i] > temperatures[0] {
				count = i
				break
			}
		}
		ans[ansBiao] = count
		ansBiao++
		temperatures = temperatures[1:]
	}
	fmt.Println(ans)
	return ans
}
func main() {
	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	dailyTemperatures(a)
}
