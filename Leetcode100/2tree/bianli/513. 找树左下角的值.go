package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 513. 找树左下角的值 8min
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 16:59
 */

func findBottomLeftValue(root *TreeNode) int {
	//层序遍历
	queue := make([]*TreeNode, 0)
	Nqueue := make([]*TreeNode, 0)
	queue = append(queue, root)
	Nqueue = nil
	for queue != nil {
		// 	添加下一层
		for _, node := range queue {
			if node.Left != nil {
				Nqueue = append(Nqueue, node.Left)
			}
			if node.Right != nil {
				Nqueue = append(Nqueue, node.Right)
			}
		}
		if Nqueue == nil {
			break
		}
		queue = Nqueue
		Nqueue = nil
	}
	return queue[0].Val
}
func main() {

}
