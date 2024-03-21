package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	x := make(map[int]int)
	for _, val := range nums {
		x[val]++
	}
	var flag int
	for k, v := range x {
		if v == 1 {
			flag = k
			return k
		}
	}
	return flag
}
func main() {
	a := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(a))
	//print(singleNumber(a))
}
