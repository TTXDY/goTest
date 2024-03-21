package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var baseVal int = root.Val
	// 先序遍历 发现和baseVal 不同说明不是单值树
	var preOrder func(node *TreeNode)
	var flag bool = true
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val != baseVal {
			flag = false
			return
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return flag
}

func main() {

}
