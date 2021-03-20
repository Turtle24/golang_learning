package main

import "fmt"

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	centre *Point
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	p1 := &Point{x: 1}
	c1 := Circle{4.56, p1}
	fmt.Println(c1.centre.x)
}
