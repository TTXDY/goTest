package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:[287] 寻找重复数  Un solved
* @ Author: Jay
* @ Date: 2024/1/28
* @ Time: 19:08
 */

import "fmt"

func findDuplicate(nums []int) int {
	length := len(nums)
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	sumN := 0
	for j := 1; j < length; j++ {
		sumN = sumN + j
	}
	fmt.Println(sum, sumN)
	return sum - sumN
}

func main() {
	findDuplicate([]int{1, 2, 3, 4, 2})
}
