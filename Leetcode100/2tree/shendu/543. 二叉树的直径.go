package main

import (
	"sort"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 543. 二叉树的直径
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 15:51
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 官方
func diameterOfBinaryTree2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 下面 +1 后，对于叶子节点就刚好是 0
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}

func deeP(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(deeP(node.Left), deeP(node.Right)) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var preOrder func(node *TreeNode)
	deepList := make([]int, 0)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		deepList = append(deepList, deeP(node.Left)+deeP(node.Right))
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	sort.Ints(deepList)
	return deepList[len(deepList)-1]
}
func main() {

}
