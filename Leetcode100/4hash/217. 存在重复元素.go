package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 217. 存在重复元素 5min
* @ Author: Jay
* @ Date: 2024/1/23
* @ Time: 14:09
 */

func containsDuplicate(nums []int) bool {
	hashTable := make(map[int]int)
	for _, x := range nums {
		_, ok := hashTable[x]
		if ok {
			return true
		}
		hashTable[x] = 0
	}
	return false
}

func main() {

}
