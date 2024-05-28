package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func task_24() {
	a := NewPoint(21, 43)
	b := &Point{254, 3}
	fmt.Println(a.Distance(b))
}
