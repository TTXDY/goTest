package main

func deep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := deep(node.Left) + 1
	right := deep(node.Right) + 1
	return max(left, right)
}

func diameterOfBinaryTree(root *TreeNode) int {
	left := deep(root.Left)
	right := deep(root.Right)
	return left + right
}

func main() {

}
