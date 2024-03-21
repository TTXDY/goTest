package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 1325. 删除给定值的叶子节点  18min
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 14:05
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 检查是否是叶子结点 并且值等于target
func checkYe(node *TreeNode, target int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Val == target {
		return true
	}
	return false
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 先序遍历 直到整棵树的叶子节点不等于target
	//var resTree *TreeNode = nil
	var flag bool = false
	// 先序遍历
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 是叶子节点 并且值等于target 那么删除
		if checkYe(node.Left, target) {
			node.Left = nil
			flag = true
		}
		if checkYe(node.Right, target) {
			node.Right = nil
			flag = true
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	if root.Left == nil && root.Right == nil {
		return nil
	}
	for flag == true {
		flag = false
		preOrder(root)
	}
	return root
}
func main() {

}
