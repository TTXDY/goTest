package main

import "sort"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 230. 二叉搜索树中第K小的元素 3min
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 17:15
 */

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	valList := make([]int, 0)
	valList = nil
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		valList = append(valList, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	sort.Ints(valList)
	return valList[k-1]
}

func main() {

}
