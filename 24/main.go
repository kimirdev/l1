package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) DistanceToAnotherPoint(p2 *Point) float64 {
	// sqrt((x2-x1)^2+(y2-y1)^2)
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(-4, -3)

	fmt.Println(p1.DistanceToAnotherPoint(p2))
}
