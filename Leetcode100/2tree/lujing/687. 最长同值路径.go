package main

import "sort"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 687. 最长同值路径
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 16:13
 */

func longestUnivaluePath(root *TreeNode) int {
	var l, r int = 0, 0
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	if root.Left != nil {
		if root.Val == root.Left.Val {
			l = longestUnivaluePath(root.Left) + 1
		}
	}
	if root.Right != nil {
		if root.Val == root.Right.Val {
			r = longestUnivaluePath(root.Right) + 1
		}
	}
	return l + r
}
func deeP(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(deeP(node.Left), deeP(node.Right)) + 1
}

func checkVal(x *TreeNode, y *TreeNode) bool {
	if x == nil || y == nil {
		return false
	}
	baseVal := x.Val
	var flag bool = true
	var flag1, flag2 bool = true, true
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val != baseVal {
			flag = false
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(x)
	flag1 = flag
	flag = true
	preOrder(y)
	flag2 = flag
	return flag2 && flag1
}

func longestUnivaluePath2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	var preOrder func(node *TreeNode)
	deepList := make([]int, 0)
	deepList = nil
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if checkVal(node.Left, node.Right) {
			deepList = append(deepList, deeP(node.Left)+deeP(node.Right))
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	temp := 0
	if deepList != nil {
		sort.Ints(deepList)
		temp = deepList[len(deepList)-1]
	}
	return temp
}

func main() {

}
