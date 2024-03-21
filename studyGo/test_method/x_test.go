package test_method

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @ Tool：IntelliJ IDEA
 * @ Author：Jay
 * @ Date：2024-01-31-16:11
 * @ Version：1.0
 * @ Description：
 */
// A Path is a journey connecting the points with straight lines.
type Point struct{ X, Y float64 }
type Path []Point

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func Test_method(t *testing.T) {
	A := Point{0, 0}
	B := Point{3, 4}
	C := Point{5, 12}
	x := Path{A, B, C}
	sum := x.Distance()
	fmt.Println(A.Distance(B))
	fmt.Println(Distance(B, C))
	fmt.Println(sum)

}
