package main

import "sort"

func findSecondMinimumValue(root *TreeNode) int {
	VlaMap := make(map[int]int)
	valList := make([]int, 0)
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		//VlaMap = append(VlaMap, node.Val)
		VlaMap[node.Val] = node.Val
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	for _, v := range VlaMap {
		valList = append(valList, v)
	}
	sort.Ints(valList)
	if valList[0] == valList[len(valList)-1] {
		return -1
	}
	return valList[1]
}

func main() {

}
