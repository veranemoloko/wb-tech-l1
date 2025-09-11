package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p1 Point) Distance(p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)

}

func main() {
	p1 := NewPoint(1.2, 2.2)
	p2 := NewPoint(14.2, -6.3)

	fmt.Printf("Distance: %.2f\n", p1.Distance(p2))
}
