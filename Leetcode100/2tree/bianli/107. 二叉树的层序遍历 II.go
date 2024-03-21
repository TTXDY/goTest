package main

import "fmt"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:
* @ Author: Jay
* @ Date: 2024/1/11
* @ Time: 12:18
 */

func reverse(num [][]int) [][]int {
	length := len(num)
	for i := 0; i < length/2; i++ {
		num[i], num[length-1-i] = num[length-1-i], num[i]
	}
	fmt.Println(num)
	return num
}
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([]int, 0)
	temp = append(temp, root.Val)
	allNodeVal := make([][]int, 0)
	allNodeVal = append(allNodeVal, temp)
	queue := make([]*TreeNode, 0)
	Newqueue := make([]*TreeNode, 0)
	NewqueueVal := make([]int, 0)
	NewqueueVal = nil
	fmt.Println(NewqueueVal)
	allNode := make([][]*TreeNode, 0)
	queue = append(queue, root)
	allNode = append(allNode, queue)
	for queue != nil {
		// 将队列中的每个节点的左右节点放进新的队列
		for _, node := range queue {
			if node.Left != nil {
				Newqueue = append(Newqueue, node.Left)
				NewqueueVal = append(NewqueueVal, node.Left.Val)

			}
			if node.Right != nil {
				Newqueue = append(Newqueue, node.Right)
				NewqueueVal = append(NewqueueVal, node.Right.Val)
			}
		}
		allNode = append(allNode, Newqueue)
		if NewqueueVal != nil {
			allNodeVal = append(allNodeVal, NewqueueVal)
		}
		queue = Newqueue
		NewqueueVal = nil
		Newqueue = nil
	}
	fmt.Println(allNodeVal)
	return reverse(allNodeVal)
}

func main() {
	tree := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(levelOrderBottom(&tree))
}
