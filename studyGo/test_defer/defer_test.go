package test_defer_test

import (
	"fmt"
	"os"
	"testing"
)

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: defer语句经常被用于处理成对的操作,如打开、关闭、连接、断开连接、加锁、释放锁。
	通过defer机制,不论函数逻辑多复杂,都能保证在任何执行路径下,资源被释放。
	释放资源的defer应该直接跟在请求资源的语句后。
* @ Author: Jay
* @ Date: 2024/1/31
* @ Time: 14:37
*/

func Test_defer2(t *testing.T) {
	filename := "/Users/jay/IdeaProjects/Golang/goTest1/studyGo/studyGo.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("heelo")
	defer fmt.Println("world")
	defer f.Close()
}
