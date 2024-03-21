package main

import (
	"fmt"
	"sort"
)

func verticalTraversal(root *TreeNode) [][]int {
	allValue := make([][]int, 0)
	length := 0
	treeMap := make(map[*TreeNode]int)
	treeMap[root] = 0
	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			treeMap[node.Left] = treeMap[node] - 1
		}
		if node.Right != nil {
			treeMap[node.Right] = treeMap[node] + 1
		}
		pre(node.Left)
		pre(node.Right)
		length++
	}
	pre(root)
	fmt.Println("root' length is ", length)
	li := make([]int, 0)
	for _, i := range treeMap {
		li = append(li, i)
	}
	sort.Ints(li)
	Min := li[0]
	Max := li[len(li)-1]
	temp := make([]int, 0)
	for i := Min; i <= Max; i++ {
		for node, flag := range treeMap {
			if flag == i {
				temp = append(temp, node.Val)
			}
		}
		allValue = append(allValue, temp)
		temp = nil
	}

	return allValue
}

func verticalTraversal3(root *TreeNode) [][]int {
	//allValue := make([][]int, 0)
	length := 0
	treeMap := make(map[*TreeNode]int)
	treeMap[root] = 0
	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			treeMap[node.Left] = treeMap[node] - 1
		}
		if node.Right != nil {
			treeMap[node.Right] = treeMap[node] + 1
		}
		pre(node.Left)
		pre(node.Right)
		length++
	}
	pre(root)
	fmt.Println("root' length is ", length)
	li := make([]int, 0)
	for _, i := range treeMap {
		li = append(li, i)
	}
	sort.Ints(li)
	Min := li[0]
	var theMin int
	if Min <= 0 {
		theMin = -Min
	}
	Max := li[len(li)-1]
	temp := make([][]int, Max+1+theMin)
	var pre2 func(node *TreeNode)
	pre2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		for i := Min; i <= Max; i++ {
			if treeMap[node] == i {
				temp[i+theMin] = append(temp[i+theMin], node.Val)
			}
		}
		pre2(node.Left)
		pre2(node.Right)
		//length++
	}
	pre2(root)

	return temp
}
func main() {

}
