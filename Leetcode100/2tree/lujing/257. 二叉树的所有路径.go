package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 257. 二叉树的所有路径   错误
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 14:37
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkYe(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func binaryTreePathsRight(root *TreeNode) []string {
	var res []string
	buildPath(root, "", &res)
	return res
}

func buildPath(root *TreeNode, pathStr string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		pathStr += strconv.Itoa(root.Val)
		*res = append(*res, pathStr)
		return
	}
	pathStr += strconv.Itoa(root.Val) + "->"
	buildPath(root.Left, pathStr, res)
	buildPath(root.Right, pathStr, res)
}

func binaryTreePaths2(root *TreeNode) []string {
	str := make([]string, 0)
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		str = append(str, strconv.Itoa(root.Val))
		return str
	}

	// 先序遍历
	var preOrder func(node *TreeNode)
	var flag bool = false
	listVal := make([]int, 0)
	alllistVal := make([][]int, 0)
	preOrder = func(node *TreeNode) {
		if node == nil {
			//listVal = nil
			return
		}
		if node.Val != math.MaxInt32 {
			listVal = append(listVal, node.Val)
		}
		if checkYe(node.Left) && node.Val != math.MaxInt32 {
			listVal = append(listVal, node.Left.Val)

			alllistVal = append(alllistVal, listVal)
			copy(listVal, listVal[:len(listVal)-1])
			if node.Right != nil {
				listVal = append(listVal, node.Right.Val)
				alllistVal = append(alllistVal, listVal)
				node.Right.Val = math.MaxInt32
			}
			node.Val = math.MaxInt32
			node.Left.Val = math.MaxInt32
			listVal = nil
			flag = true
		}
		if checkYe(node.Right) && node.Val != math.MaxInt32 {
			listVal = append(listVal, node.Right.Val)
			alllistVal = append(alllistVal, listVal)
			//listVal = listVal[:len(listVal)-1]
			copy(listVal, listVal[:len(listVal)-1])
			if node.Left != nil {
				listVal = append(listVal, node.Left.Val)
				alllistVal = append(alllistVal, listVal)
				node.Left.Val = math.MaxInt32
			}
			node.Val = math.MaxInt32
			node.Right.Val = math.MaxInt32
			listVal = nil
			flag = true
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	for flag == true {
		flag = false
		preOrder(root)
	}
	fmt.Println(alllistVal)
	return nil
}

func binaryTreePaths(root *TreeNode) []string {
	listVal := make([]int, 0)
	alllistVal := make([][]int, 0)

	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if checkYe(node) == true && node.Val != math.MaxInt32 {
			listVal = append(listVal, node.Val)
			alllistVal = append(alllistVal, listVal)
			node.Val = math.MaxInt32
			return
		}
		listVal = append(listVal, node.Val)
		pre(node.Left)
		pre(node.Right)
		listVal = nil
	}
	pre(root)
	fmt.Println(alllistVal)
	return nil
}
func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	binaryTreePaths(&tree)
}
