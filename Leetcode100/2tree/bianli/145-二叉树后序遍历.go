package main

func postorderTraversal(root *TreeNode) []int {
	valueList := make([]int, 0)
	var post func(node *TreeNode)
	post = func(node *TreeNode) {
		if node == nil {
			return
		}
		post(node.Left)
		post(node.Right)
		valueList = append(valueList, node.Val)
	}
	post(root)
	return valueList
}
