package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 566. 重塑矩阵
* @ Author: Jay
* @ Date: 2024/1/25
* @ Time: 11:31
 */

func matrixReshape(mat [][]int, r int, c int) [][]int {
	hangLength := len(mat)
	lieLength := len(mat[0])
	if hangLength*lieLength != r*c {
		return mat
	}
	valueList := make([]int, 0)
	for _, v := range mat {
		for _, vv := range v {
			valueList = append(valueList, vv)
		}
	}
	res := make([][]int, 0)
	k := 0
	for i := 0; i < r; i++ {
		var temp []int
		for j := 0; j < c; j++ {
			temp = append(temp, valueList[k])
			k++
		}
		res = append(res, temp)
	}
	return res
}

func main() {

}
