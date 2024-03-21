package main

func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var inorder func(node *TreeNode, node2 *TreeNode) *TreeNode
	inorder = func(node *TreeNode, node2 *TreeNode) *TreeNode {
		if node == nil && node2 != nil {
			node = node2
			node.Right = nil
			node.Left = nil
			return node
		}
		if node2 == nil && node != nil {
			node2 = node
			node2.Right = nil
			node2.Left = nil
			return node
		}
		if node != nil && node2 != nil {
			node.Val = node.Val + node2.Val
			return node
		}
		inorder(node.Left, node2.Left)
		inorder(node.Right, node2.Right)
		return node
	}
	return inorder(root1, root2)
}
func mergeTrees3(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees3(t1.Left, t2.Left)
	t1.Right = mergeTrees3(t1.Right, t2.Right)
	return t1
}
func main() {

}
