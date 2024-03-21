package main

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var pre func(node *Node)
	valList := make([]int, 0)
	pre = func(node *Node) {
		if node.Children == nil {
			return
		}
		valList = append(valList, node.Val)
		for i := 0; i < len(node.Children); i++ {
			pre(node.Children[i])
		}
	}
	pre(root)
	return valList
}

func main() {

}
