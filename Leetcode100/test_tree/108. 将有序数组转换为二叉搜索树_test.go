package test_tree

import (
	"fmt"
	"math"
	"testing"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 108. 将有序数组转换为二叉搜索树
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 20:22
 */

/**
 * Definition for a binary 2tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(treeNode *TreeNode) {
	value := make([]int, 0)
	var inorder2 func(node *TreeNode)
	inorder2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder2(node.Left)
		value = append(value, node.Val)
		inorder2(node.Right)
	}
	inorder2(treeNode)
	fmt.Println(value)
}
func constructTree(root *TreeNode, deep int) *TreeNode {
	if deep == 1 {
		return root
	}
	theNode := TreeNode{
		Val:   math.MaxInt32,
		Left:  nil,
		Right: nil,
	}
	root.Left = &theNode
	theNode2 := TreeNode{
		Val:   math.MaxInt32,
		Left:  nil,
		Right: nil,
	}
	root.Right = &theNode2
	constructTree(root.Left, deep-1)
	constructTree(root.Right, deep-1)
	return root
}
func sortedArrayToBST(nums []int) *TreeNode {
	//层序遍历建立二叉树
	nodeNum := len(nums)
	if nodeNum == 0 {
		return nil
	}
	var deep float64 = 0
	for math.Pow(2.0, deep)-1 < float64(nodeNum/2) {
		deep = deep + 1.0
	}
	root := &TreeNode{
		Val:   99,
		Left:  nil,
		Right: nil,
	}
	root = constructTree(root, int(deep))
	PrintTree(root)
	var inorder func(node *TreeNode)
	i := 0
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if i < nodeNum {
			node.Val = nums[i]
			i++
		}
		inorder(node.Right)
	}
	inorder(root)
	PrintTree(root)

	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		pre(node.Left)
		if node.Left != nil && node.Left.Val == math.MaxInt32 {
			node.Left = nil
		}
		if node.Right != nil && node.Right.Val == math.MaxInt32 {
			node.Right = nil
		}
		pre(node.Right)

	}
	pre(root)
	PrintTree(root)
	return root
}

func TestN(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	sortedArrayToBST(a)
}
