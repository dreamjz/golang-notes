package chpater_10

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

type NamedPoint2 struct {
	Point
	name string
}

func (np *NamedPoint2) Abs() float64 {
	return np.Point.Abs() * 100
}

func mh() {
	n := &NamedPoint{Point{3, 4}, "Node"}
	fmt.Println(n.Abs())

	n2 := &NamedPoint2{Point{3, 4}, "Node2"}
	fmt.Println(n2.Abs())
	fmt.Println(n2.Point.Abs())
}
