/*
append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
*/
package main

import "fmt"

func main() {
	var numbers []int
	printSlice1(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice1(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice1(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice1(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice1(numbers1)
}

// printSlice1
func printSlice1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
