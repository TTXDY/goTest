package main

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	valList := make([]int, 0)
	var post func(node *Node)
	post = func(node *Node) {
		if node.Children == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			post(node.Children[i])
		}
		valList = append(valList, node.Val)
	}
	post(root)
	return valList
}

func main() {

}
