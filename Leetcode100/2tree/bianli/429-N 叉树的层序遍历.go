package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	list := make([]int, 0)
	allList := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		newQueue := make([]*Node, 0)
		// 把队列中的所有节点的值取出来存放在list中
		for _, v := range queue {
			list = append(list, v.Val)
		}
		allList = append(allList, list)
		for _, node := range queue {
			for _, child := range node.Children {
				newQueue = append(newQueue, child)
			}
		}
		// 删除队列中所有元素, 添加新的节点
		queue = newQueue
		list = nil
	}
	return allList
}

// 层序遍历 迭代法不好使
func levelOrder2(root *Node) [][]int {
	allList := make([][]int, 0)
	var level func(node *Node)
	level = func(node *Node) {
		var getValue func(node2 *Node) []int
		getValue = func(node2 *Node) []int {
			valList := make([]int, 0)
			for _, v := range node2.Children {
				valList = append(valList, v.Val)
			}
			return valList
		}

		if node.Children == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			allList = append(allList, getValue(node.Children[i]))
		}

	}
	level(root)
	return allList

}

func main() {

}
