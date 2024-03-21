package main

import (
	"fmt"
	"goTest1/Init/13-fengzhuang/test"
)
import "goTest1/Init/13-fengzhuang/test2"

func main() {
	h := test.Human{}
	h.Eat()
	h.Walk()

	// 初始化变量的两种方法，直接定义 和  使用var 声明再设置
	//s := Superman{
	//	Human{Sex: 0, name: "li4"},
	//	10,
	//}
	var s test.Superman
	s.Sex = 0
	//s.name = "jay"
	s.Level = 90
	s.Eat()
	s.Fly()
	println("+++++++++++++_____++++++++++++++++")

	p := test2.Point{X: 0, Y: 0.0}
	var q test2.Point
	dis := p.Distance(test2.Point{X: 3, Y: 4.0})
	println("distance1 of Point q and p is ", dis)

	q.X = 5
	q.Y = 12
	dis2 := p.Distance(q)
	println("distance2 of Point q and p is ", dis2)

	//println("+++++++++++++_____++++++++++++++++")

	k := test2.Point{
		X: 8,
		Y: 15,
	}
	//方法值(相当于C语言的函数地址,函数指针)
	disMethod := p.Distance
	dis3 := disMethod(k)
	println("distance3 of Point q and p is ", dis3)

	//方法表达式
	disMethod2 := test2.Point.Distance
	dis4 := disMethod2(p, k)
	println("distance4 of Point q and p is ", dis4)
	fmt.Printf("%T\n", dis4) //%T表示打出数据类型 ,这个必须放在Printf使用

	//方法表达式2
	disMethod3 := (*test2.Point).Distance
	dis5 := disMethod3(&p, k)
	println("distance of Point q and p is ", dis5)
	fmt.Printf("%T\n", dis5) //%T表示打出数据类型 ,这个必须放在Printf使用

	println("+====================================+")

	points := test2.Path{
		{100, 100},
		{120, 120},
	}

	anotherPoint := test2.Point{X: 5, Y: 5}

	points.TranslateBy(anotherPoint, false)

	fmt.Println("------------------")

	points1 := test2.Path{
		{100, 100},
		{120, 120},
	}

	points1.TranslateBy(anotherPoint, true)

}
