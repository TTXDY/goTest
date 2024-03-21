package main

import (
	"fmt"
)

// 官方
func isBalanced0(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 牛的
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func returnDeep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var Leftdeep, Rightdeep int

	Leftdeep = returnDeep(node.Left) + 1
	Rightdeep = returnDeep(node.Right) + 1
	return max(Leftdeep, Rightdeep)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 返回深度
	var flag bool
	flag = true
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		// 非常重要！！！
		if node == nil || flag == false {
			return
		}
		a := returnDeep(node.Left) - returnDeep(node.Right)
		if a == 1 || a == 0 || a == -1 {
			flag = true
			fmt.Println(a, flag)
		} else {
			flag = false
			fmt.Println(a, flag)
			return
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return flag
}

// 错误
func isBalanced2(root *TreeNode) bool {
	// 返回深度
	var flag bool
	flag = true
	if root == nil {
		return true
	}
	if flag == false {
		return false
	}
	// 非常重要！！！
	if root == nil {
		return true
	}
	a := returnDeep(root.Left) - returnDeep(root.Right)
	if a == 1 || a == 0 || a == -1 {
		flag = true
		fmt.Println(a, flag)
	} else {
		flag = false
		fmt.Println(a, flag)
		return false
	}
	isBalanced2(root.Left)
	isBalanced2(root.Right)

	return flag
}

func main() {
	// 返回深度
	var returnDeep func(node *TreeNode) int
	var Leftdeep, Rightdeep int

	returnDeep = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		Leftdeep = returnDeep(node.Left) + 1
		Rightdeep = returnDeep(node.Right) + 1
		return max(Leftdeep, Rightdeep)
	}
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil},
		},
		Right: nil,
	}
	fmt.Println(isBalanced(&tree))
}
