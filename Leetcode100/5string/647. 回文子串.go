package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 647. 回文子串 13min
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 11:00
 */

func judgeHUIWEN(n []rune) bool {
	length := len(n)
	for i := 0; i < length/2+1; i++ {
		if n[i] != n[length-1-i] {
			return false
		}
	}
	return true
}

func countSubstrings(s string) int {
	if len(s) == 1 {
		return 1
	}
	arrS := []rune(s)
	count := 0
	window := 1
	i := 0
	for window < len(s) {
		for ; i < len(s)+1-window; i++ {
			temp := arrS[i : i+window]
			fmt.Println(temp)
			if judgeHUIWEN(temp) {
				count++
			}
		}
		i = 0
		window++
	}

	return count
}

func main() {

}
