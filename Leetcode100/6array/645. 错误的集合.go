package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 645. 错误的集合  Unsolved
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 14:07
 */

func findErrorNums(nums []int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	ans := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := cnt[i]; c == 2 {
			ans[0] = i
		} else if c == 0 {
			ans[1] = i
		}
	}
	return ans
}

func main() {

}
