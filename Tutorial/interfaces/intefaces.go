package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shapes) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	r1 := Rect{3, 5}
	shapes := []shapes{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
