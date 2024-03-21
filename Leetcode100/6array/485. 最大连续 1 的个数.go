package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 485. 最大连续 1 的个数
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 13:28
 */

func findMaxConsecutiveOnes(nums []int) int {
	var count int = 0
	var theMax int = 0
	for _, num := range nums {
		if num == 1 {
			count++
			if count > theMax {
				theMax = count
			}
		} else {
			count = 0
		}
	}
	return theMax
}

func main() {

}
