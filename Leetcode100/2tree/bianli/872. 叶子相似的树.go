package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 872. 叶子相似的树  20min
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 13:11
 */

// 判断是不是叶子节点
func checkYe(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func checkNum(numP []int, numQ []int) bool {
	if len(numQ) != len(numP) {
		return false
	}
	for i, v := range numP {
		if v != numQ[i] {
			return false
		}
	}
	return true
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 先序遍历 分别记录root1和root2的叶子节点序列对，比较是否相同
	// 先序遍历
	var preOrder func(node *TreeNode)
	valList := make([]int, 0)
	firstYe := make([]int, 0)
	sencondYe := make([]int, 0)
	firstYe, sencondYe = nil, nil
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if checkYe(node) {
			valList = append(valList, node.Val)
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root1)
	firstYe = valList
	valList = nil
	preOrder(root2)
	sencondYe = valList
	fmt.Println(firstYe, sencondYe)
	return checkNum(firstYe, sencondYe)
}
func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	tree2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	leafSimilar(&tree, &tree2)
}
