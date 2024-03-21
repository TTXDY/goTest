package main

import "fmt"

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-02-28-21:27
 * @ Version：1.0
 * @ Description：11. 盛最多水的容器 暴力法、双指针
 */

// 暴力法
func maxArea(height []int) int {
	Max, sum := 0, 0
	legth := len(height)
	for i := 0; i < legth; i++ {
		for j := i; j < legth; j++ {
			sum = min(height[i], height[j]) * (j - i)
			fmt.Println(height[i], height[j], sum)
			if sum > Max {
				Max = sum
			}
		}
	}
	return Max
}
