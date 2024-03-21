package main

import (
	"fmt"
	"sort"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 594. 最长和谐子序列 40min
* @ Author: Jay
* @ Date: 2024/1/23
* @ Time: 14:21
 */
func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}
func findLHS(nums []int) int {
	hashTable := make(map[int]int)
	for _, x := range nums {
		_, ok := hashTable[x]
		if ok {
			hashTable[x]++
		} else {
			hashTable[x] = 1
		}

	}
	fmt.Println(hashTable)
	theMax := 0
	keyList := make([]int, 0)

	for k, _ := range hashTable {
		keyList = append(keyList, k)
	}
	fmt.Println(keyList)
	sort.Ints(keyList)
	for i := 0; i < len(keyList)-1; i++ {
		if theMax < hashTable[keyList[i]]+hashTable[keyList[i+1]] && abs(keyList[i]-keyList[i+1]) == 1 {
			theMax = hashTable[keyList[i]] + hashTable[keyList[i+1]]
		}
	}
	return theMax
}

func main() {
	a := []int{0, 3, 1, 3, 3, 3, 0, 1, 0, 2, 0, 3, 1, 3, -3, 2, 0, 3, 1, 2, 2, -3, 2, 2, 3, 3}
	fmt.Println(findLHS(a))
}
