package main

func printArray(parr []int, lens int) {
	for i := 0; i < lens; i++ {
		println("index: ", i, "value: ", parr[i])
	}
	parr[0] = 100
}

func main() {
	arr := []int{1, 3, 4, 5}
	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}
	println(`++++++======+++++`)

	for index, item := range arr {
		print(index, " ", item, "\n")
	}
	printArray(arr, len(arr))
	for index, item := range arr {
		print(index, " ", item, "\n")
	}
}
