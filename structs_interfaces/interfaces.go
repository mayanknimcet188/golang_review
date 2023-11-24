package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Point struct {
	x, y float64
}

func distance(a Point, b Point) float64 {
	x_dist := b.x - a.x
	y_dist := b.y - a.y
	return math.Sqrt(x_dist*x_dist + y_dist*y_dist)
}

type Rectangle struct {
	a, b Point
}

func (r *Rectangle) area() float64 {
	c := Point{x: r.b.x, y: r.a.y}
	d := Point{x: r.a.x, y: r.b.y}

	length := distance(r.a, c)
	width := distance(r.a, d)
	return length * width
}

// interfaces contain "method sets" instead of fields
// A type must implement the list of methods in the method set in order to
// implement the inteface
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{r: 5}
	a := Point{x: 0, y: 0}
	b := Point{x: 10, y: 10}
	r := Rectangle{a: a, b: b}
	fmt.Println(totalArea(&c, &r))
}
