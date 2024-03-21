package main

import "sort"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 378. 有序矩阵中第 K 小的元素
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 13:52
 */

func kthSmallest(matrix [][]int, k int) int {
	valueList := make([]int, 0)
	for _, row := range matrix {
		for _, v := range row {
			valueList = append(valueList, v)
		}
	}
	sort.Ints(valueList)
	return valueList[k-1]
}

func main() {

}
