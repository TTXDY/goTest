package main

func deep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	a := deep(node.Left) + 1
	b := deep(node.Right) + 1
	return max(a, b)
}
func maxDepth(root *TreeNode) int {
	x1 := deep(root.Right)
	x2 := deep(root.Left)
	return max(x1, x2)
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	//var a int

}
