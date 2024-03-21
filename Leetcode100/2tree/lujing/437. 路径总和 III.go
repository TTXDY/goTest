package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 437. 路径总和 III
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 15:26
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rtrnNum(node *TreeNode, target int) (ans int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && target == node.Val {
		ans = ans + 1
	}
	rtrnNum(node.Left, target-node.Val)
	rtrnNum(node.Right, target-node.Val)
	return ans
}

func pathSum(root *TreeNode, targetSum int) int {
	var pre func(node *TreeNode)
	var ans int = 0
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		if targetSum == node.Val {
			temp := node.Val
			ans = ans + 1
			targetSum = targetSum - temp
		}
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return ans
}

func main() {

}
