package main

import "fmt"

func isHuiwen(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	queue := make([]int, 0)
	for _, v := range s {
		queue = append(queue, int(v))
	}
	length := len(queue)
	fmt.Println(queue)
	for q := 0; q < length; q++ {
		if queue[q] != queue[length-1-q] {
			return false
		}
	}
	return true
}
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	left := 0
	right := 0
	length := 0
	longString := ""
	Slength := len(s)
	queue := make([]int, 0)
	for _, v := range s {
		queue = append(queue, int(v))
	}
	i := 0
	for i <= Slength {
		if right < Slength {
			newS := s[left : right+1]
			if isHuiwen(newS) {
				if right-left+1 > length {
					longString = newS
					length = right - left + 1
				}
			}
		}
		if right == len(s) {
			s = s[1:]
			Slength = len(s)
			right = 0
			i = -1
		}
		right++
		i++
	}
	return longString
}

func main() {
	a := "babad"
	fmt.Println(longestPalindrome(a))
	//fmt.Println(isHuiwen(a[0:4]))
}
