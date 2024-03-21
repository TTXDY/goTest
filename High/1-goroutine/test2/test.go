package main

import (
	"fmt"
	"runtime"
)

func Ado() {
	defer fmt.Println("B.defer")
	runtime.Goexit() // 终止当前 goroutine, import "runtime"
	fmt.Println("B") // 不会执行
}

func Do() {
	defer fmt.Println("A.defer")

	Ado()

	fmt.Println("A") // 不会执行
}

func main() {
	go Do()

	//死循环，目的不让主goroutine结束
	for {
	}
}
