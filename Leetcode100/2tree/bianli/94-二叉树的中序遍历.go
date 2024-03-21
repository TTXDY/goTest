package main

func inorderTraversal(root *TreeNode) []int {
	value := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		value = append(value, node.Val)
		inorder(node.Right)

	}
	inorder(root)
	return value
}

func main() {

}
