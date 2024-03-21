package main

func preorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return list
}
