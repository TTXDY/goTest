package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 501. 二叉搜索树中的众数
* @ Author: Jay
* @ Date: 2024/1/13
* @ Time: 14:24
 */

/**
 * Definition for a binary 2tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var inorder func(node *TreeNode)
	valList := make([]int, 0)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		valList = append(valList, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	fmt.Println(valList)
	var count int = 1
	//res := make([]int, 0)
	//res = append(res, 1)
	temp := 1
	zhongshu := make([]int, 0)
	for i := 0; i < len(valList)-1; i++ {
		if valList[i] == valList[i+1] {
			count++
			if count > temp {
				zhongshu = nil
				zhongshu = append(zhongshu, valList[i])
				temp = count
			} else if count == temp {
				zhongshu = append(zhongshu, valList[i])
			}
		} else {
			count = 1
		}
	}
	if temp == 1 {
		return valList
	}
	return zhongshu
}

func main() {

}
