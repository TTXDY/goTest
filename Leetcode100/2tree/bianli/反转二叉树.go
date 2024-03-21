package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root

	}
	var temp *TreeNode
	temp = root.Right
	root.Right = root.Left
	root.Left = temp
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}
func main() {

}
