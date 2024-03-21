package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	行 := 10

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
	fmt.Println("type: ", reflect.TypeOf(行))
	fmt.Println("value: ", reflect.ValueOf(行))
}
