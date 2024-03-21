package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
}

func Hello() {
	i := 0
	//main goroutine 循环打印
	for {
		i++
		fmt.Printf("hello goroutine: i = %d\n", i)
		time.Sleep(5 * time.Second) //延时5s
	}
}

func main() {
	//创建一个 goroutine，启动另外一个任务
	go newTask()
	go Hello()

	i := 0
	//main goroutine 循环打印
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(5 * time.Second) //延时5s
	}
	fmt.Println("112233")

}
