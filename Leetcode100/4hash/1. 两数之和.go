package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 1. 两数之和 40min
* @ Author: Jay
* @ Date: 2024/1/23
* @ Time: 13:25
 */

func twoSum(nums []int, target int) []int {
	resultList := make([]int, 0)
	numsMap := make(map[int]int)
	for i, num := range nums {
		p, ok := numsMap[target-num]
		if ok {
			resultList = append(resultList, p)
			resultList = append(resultList, i)
			return resultList
		}
		numsMap[num] = i
	}
	return nil
}

func main() {

}
