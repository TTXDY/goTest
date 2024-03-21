package main

import "fmt"

func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("It is not ok for type 5string")
		return ""
	}
	fmt.Println("The value is ", value)

	return value
}

func main() {
	//      str := "123"
	//      funcName(str)
	//var a interface{}
	var a int = 100
	var b string = "123"
	funcName(a)
	funcName(b)
}
