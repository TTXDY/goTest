package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 205. 同构字符串
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 10:26
 */

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := make(map[rune]rune)
	t2s := make(map[rune]rune)
	arrS := []rune(s)
	arrT := []rune(t)
	for i := 0; i < len(arrS); i++ {
		_, ok1 := s2t[arrS[i]]
		_, ok2 := t2s[arrT[i]]
		if ok1 == false && ok2 == true || ok1 == true && ok2 == false {
			return false
		}
		if ok1 && ok2 {
			if s2t[arrS[i]] != arrT[i] {
				return false
			}
		}
		s2t[arrS[i]] = arrT[i]
		t2s[arrT[i]] = arrS[i]
	}
	return true
}

func main() {
	s := "bbbaaaba"
	t := "aaabbbba"
	fmt.Println(isIsomorphic(s, t))
}
