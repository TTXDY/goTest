package main

import "fmt"

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-02-28-20:22
 * @ Version：1.0
 * @ Description：560. 和为 K 的子数组  只会用暴力解  不会最优解，但值得学习，很重要！
 */

// 暴力法
func subarraySum(nums []int, k int) int {
	var count int = 0
	length := len(nums)
	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += nums[j]
			fmt.Println("i, j, sum", i, j, sum)
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 最优解
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func main() {
	num := []int{1, 1, 1}
	print(subarraySum(num, 2))
}
