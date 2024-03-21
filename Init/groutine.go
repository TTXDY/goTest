package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int64 = 20
	var y int64 = 10
	var wg sync.WaitGroup

	wg.Add(1)
	//定义一个匿名函数，并对该函数开启协程
	go func(x, y int64) {
		z := x + y
		fmt.Println("the reuslt value:", z)
		wg.Done()
	}(x, y)
	//由于这个函数是匿名函数，所以调用方式就直接是（x,y）去调用，不用输入函数名。

	wg.Wait()

}
