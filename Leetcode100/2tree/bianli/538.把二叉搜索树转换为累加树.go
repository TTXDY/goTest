package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:  538 把二叉搜索树转换为累加树 18min
* @ Author: Jay
* @ Date: 2024/1/12
* @ Time: 19:27
 */

func doNum(num []int) []int {
	for i := 1; i < len(num); i++ {
		num[i] = num[i] + num[i-1]
	}
	return num
}
func convertBST(root *TreeNode) *TreeNode {
	reverValList := make([]int, 0)
	valList := make([]int, 0)
	var reverInorder func(node *TreeNode)
	reverInorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		reverInorder(node.Right)
		reverValList = append(reverValList, node.Val)
		reverInorder(node.Left)
	}
	reverInorder(root)
	valList = doNum(reverValList)
	var reverInorder2 func(node *TreeNode)
	reverInorder2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		reverInorder2(node.Right)
		node.Val = valList[0]
		valList = valList[1:]
		reverInorder2(node.Left)
	}
	reverInorder2(root)
	return root
}

func main() {

}
