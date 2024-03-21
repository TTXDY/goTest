package main

import (
	"fmt"
	"sort"
)

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-02-28-21:02
 * @ Version：1.0
 * @ Description：581. 最短无序连续子数组 暴力法
 */

// 先排序，然后对比原来的数组找距离
func findUnsortedSubarray(nums []int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	fmt.Println(numsCopy, nums)
	left, count := 0, 0
	flag := true

	for i := 0; i < len(nums); i++ {
		if numsCopy[i] != nums[i] && flag {
			left = i
			flag = false
		}
		if numsCopy[i] != nums[i] && !flag {
			if i-left+1 > count {
				count = i - left + 1
			}
		}
	}
	return count
}
