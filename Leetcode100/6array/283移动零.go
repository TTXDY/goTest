package main

import "fmt"

func returnShu(nums []int) int {
	var shuPoint int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			shuPoint = i
			break
		}
	}
	return shuPoint
}
func returnZero(nums []int) int {
	var zeroPoint int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroPoint = i
			break
		}
	}
	return zeroPoint
}

func exchangeNotZero(nums []int) int {
	var shuPoint int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			shuPoint = i
			break
		}
	}
	return shuPoint
}
func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			index := exchangeNotZero(nums[i:])
			var temp int
			temp = nums[i]
			nums[i] = nums[index+i]
			nums[index+i] = temp
		}
	}
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			index := exchangeNotZero(nums[i:])
			var temp int
			temp = nums[i]
			nums[i] = nums[index+i]
			nums[index+i] = temp
		}
	}
	fmt.Println(nums)
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
}
