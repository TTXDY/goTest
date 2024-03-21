package test_defer

import (
	"fmt"
	"testing"
)

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-01-31-15:19
 * @ Version：1.0
 * @ Description：
 */
func f() (result int) {
	defer func() {
		result = result - 1
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1.x = 5 2. defer x = 6 3  真正的返回
}

func f6() (x int) {
	defer func(x *int) *int {
		*x++
		return x
	}(&x)
	return 5 // 1. x = 5 // 2.defer  x =6  3. ret返回
}

// 原文链接：https://blog.csdn.net/qq_15371293/article/details/124187385
func Test_defer(t *testing.T) {
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
