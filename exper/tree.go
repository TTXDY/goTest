package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(root *TreeNode, deep int, nodeNum int) *TreeNode {
	if deep == 1 || nodeNum == 0 {
		return root
	}
	var k int = 0
	if nodeNum == 1 {
		theNode := TreeNode{
			Val:   deep,
			Left:  nil,
			Right: nil,
		}
		root.Left = &theNode
		k++
	}
	if nodeNum >= 2 {
		theNode := TreeNode{
			Val:   deep,
			Left:  nil,
			Right: nil,
		}
		root.Left = &theNode
		theNode2 := TreeNode{
			Val:   deep,
			Left:  nil,
			Right: nil,
		}
		root.Right = &theNode2
		k = k + 2
	}

	constructTree(root.Left, deep-1, nodeNum-k)
	constructTree(root.Right, deep-1, nodeNum-k)
	return root
}

func main() {
	//n := make([]int, 10)
	root := &TreeNode{
		Val:   99,
		Left:  nil,
		Right: nil,
	}
	constructTree(root, 4, 5)
	value := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		value = append(value, node.Val)
		inorder(node.Right)

	}
	inorder(root)
	fmt.Println(value)
}
