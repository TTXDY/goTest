package main

import "math"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 530. 二叉搜索树的最小绝对差 15min
* @ Author: Jay
* @ Date: 2024/1/13
* @ Time: 14:04
 */

/**
 * Definition for a binary 2tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	minD := math.MaxInt32
	value := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		value = append(value, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	for j := 0; j < len(value); j++ {
		for i := j + 1; i < len(value); i++ {
			minD = min(abs(value[i]-value[j]), minD)
		}
	}
	return minD
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -1 * a
	}
}

func main() {

}
