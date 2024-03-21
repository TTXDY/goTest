package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 242. 有效的字母异位词 5min
* @ Author: Jay
* @ Date: 2024/1/24
* @ Time: 22:09
 */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashTableS := make(map[rune]int)
	hashTableT := make(map[rune]int)
	for _, v := range s {
		_, ok := hashTableS[v]
		if !ok {
			hashTableS[v] = 0
		}
		hashTableS[v]++
	}
	for _, v := range t {
		_, ok := hashTableT[v]
		if !ok {
			hashTableT[v] = 0
		}
		hashTableT[v]++
	}
	fmt.Println(hashTableS, hashTableT)
	for k, _ := range hashTableS {
		if hashTableS[k] != hashTableT[k] {
			return false
		}
	}
	return true
}

func main() {

}
