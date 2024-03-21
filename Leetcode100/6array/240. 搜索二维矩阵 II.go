package main

import "sort"

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 240. 搜索二维矩阵 II

	sort.SearchInts!!!
		在一个整数排序的切片中搜索x，并返回索引
		由Search指定。返回值是插入x的索引(如果x是)
		不存在(它可能是len(a))。
		切片必须按升序排序。

* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 13:41
*/

func searchMatrix(matrix [][]int, target int) bool {
	for i, row := range matrix {
		index := sort.SearchInts(row, target)
		if index != len(row) {
			if matrix[i][index] == target {
				return true
			}
		}
	}
	return false
}

func main() {

}
