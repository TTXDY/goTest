package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 236. 二叉树的最近公共祖先 cuowu
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 19:55
 */

/**
 * Definition for a binary 2tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 官方
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func main() {

}
