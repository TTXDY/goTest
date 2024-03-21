package main

import (
	"fmt"
	"math"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 128. 最长连续序列
* @ Author: Jay
* @ Date: 2024/1/23
* @ Time: 15:10
 */

func longestConsecutive(nums []int) int {
	hashTable := make(map[int]int)
	theMin := math.MaxInt32
	for _, num := range nums {
		hashTable[num] = 0
		if num < theMin {
			theMin = num
		}
	}
	//fmt.Println(theMin)
	//fmt.Println(hashTable)
	countMax := 0
	count := 0
	for k, _ := range hashTable {
		var ok bool
		_, ok = hashTable[theMin]
		for ok != false {
			count++
			theMin++
			_, ok = hashTable[theMin]
		}
		if count > countMax {
			countMax = count
		}
		theMin = k
		count = 0
	}
	return countMax
}

func main() {
	a := []int{7, -9, 3, -6, 3, 5, 3, 6, -2, -5, 8, 6, -4, -6, -4, -4, 5, -9, 2, 7, 0, 0}
	//b := []int{-2, -3, -3, 7, -3, 0, 5, 0, -8, -4, -1, 2}
	fmt.Println(longestConsecutive(a))
	//fmt.Println(longestConsecutive(b))

}
