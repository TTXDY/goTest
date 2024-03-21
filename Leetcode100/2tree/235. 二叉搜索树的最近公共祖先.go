package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:235. 二叉搜索树的最近公共祖先  10min
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 15:27
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		// 分散在两边，说明祖先必然为root
		if root.Val > p.Val && root.Val < q.Val {
			return root
		}
		if root.Val > q.Val && root.Val < p.Val {
			return root
		}
		// 相等返回root
		if root.Val == q.Val || root.Val == p.Val {
			return root
		}
		//不在两边呢 说明p, q都在右边
		if root.Val < p.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}

/**
 * Definition for a binary 2tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
func main() {

}
