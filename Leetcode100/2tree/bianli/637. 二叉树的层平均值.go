package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 637. 二叉树的层平均值 18min
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 16:32
 */

// TreeNode 层序遍历 结构体不一样，有点细微差别：要做一个左右子树的判断
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countAvg(queue []*TreeNode) float64 {
	var sum int
	if len(queue) == 0 {
		return 0
	}
	for _, v := range queue {
		if v != nil {
			sum = sum + v.Val
		}
	}
	return float64(sum) / float64(len(queue))
}

func averageOfLevels(root *TreeNode) []float64 {
	// 数组记录每层的平均数
	avgDeep := make([]float64, 0)
	// 层序遍历
	queue := make([]*TreeNode, 0)
	Nqueue := make([]*TreeNode, 0)
	Nqueue = nil
	queue = append(queue, root)
	for queue != nil {
		avgDeep = append(avgDeep, countAvg(queue))
		for _, node := range queue {
			if node.Left != nil {
				Nqueue = append(Nqueue, node.Left)
			}
			if node.Right != nil {
				Nqueue = append(Nqueue, node.Right)
			}
		}
		queue = Nqueue
		Nqueue = nil
	}
	return avgDeep
}

func main() {

}
