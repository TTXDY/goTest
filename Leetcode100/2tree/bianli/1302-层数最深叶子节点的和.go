package main

func deepestLeavesSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	Nqueue := make([]*TreeNode, 0)
	temp := make([]*TreeNode, 0)
	queue = append(queue, root)

	//flag := true
	for len(queue) != 0 {
		for _, node := range queue {
			if node.Left != nil {
				Nqueue = append(Nqueue, node.Left)
			}
			if node.Right != nil {
				Nqueue = append(Nqueue, node.Right)
			}
		}
		temp = queue
		queue = Nqueue
		Nqueue = nil
	}
	sum := 0
	for _, v := range temp {
		sum += v.Val
	}
	return sum
}

func main() {

}
