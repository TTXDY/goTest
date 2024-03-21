package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 409. 最长回文串
* @ Author: Jay
* @ Date: 2024/1/24
* @ Time: 22:39
 */

func longestPalindrome(s string) int {
	sMap := make(map[rune]int)
	var count int = 0
	theMaxSingleNum := 0
	var flag, flag2 bool
	for _, x := range s {
		_, ok := sMap[x]
		if !ok {
			sMap[x] = 0
		}
		sMap[x]++
	}
	for _, v := range sMap {
		if v%2 == 0 {
			count = count + v
		} else {
			if v > 1 {
				theMaxSingleNum = theMaxSingleNum + v - 1
				flag = true
			}
			if v == 1 {
				flag2 = true
			}
		}
	}
	fmt.Println(count, theMaxSingleNum)
	if flag == true {
		return count + theMaxSingleNum + 1
	}
	if flag == false && flag2 == true {
		return count + theMaxSingleNum + 1
	}
	return count + theMaxSingleNum
}

func main() {

}
