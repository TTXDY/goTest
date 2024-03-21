package test2

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y float64
}

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(
		float64(p.X-q.X),
		p.Y-q.Y,
	)
}

// Add 加法
func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

// Sub 减法
func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

// Print 打印点坐标
func (p Point) Print() {
	fmt.Printf("{%d, %f}\n", p.X, p.Y)
}

// Path 定义一个Point切片类型
type Path []Point

func (path Path) TranslateBy(another Point, add bool) {
	var op func(p, q Point) Point
	if add == true {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i, _ := range path {
		path[i] = op(path[i], another)
		path[i].Print()
	}
}
