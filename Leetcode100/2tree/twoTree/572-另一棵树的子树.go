package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var pre func(node *TreeNode)
	var flag bool
	flag = false
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		//if reflect.DeepEqual(node, subRoot) {
		if check(node, subRoot) {
			flag = true
		}
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return flag
}
func main() {

}
