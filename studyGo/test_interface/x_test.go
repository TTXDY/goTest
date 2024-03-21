package test_interface

import "testing"

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-01-31-16:28
 * @ Version：1.0
 * @ Description：
 */
type FirstInterface interface {
	Print()
}

func Func(x FirstInterface) {}

type PatentStruct struct{}

func (pt PatentStruct) Print() {}

func Test_interface(t *testing.T) {
	var pt PatentStruct
	Func(pt)
}
