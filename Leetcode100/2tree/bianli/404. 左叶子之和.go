package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 404. 左叶子之和 3min
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 16:01
 */

func sumOfLeftLeaves(root *TreeNode) int {
	// 先序遍历 左边叶子不为空，统计之
	var preOrder func(node *TreeNode)
	var sum int = 0
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return sum
}
func main() {

}
