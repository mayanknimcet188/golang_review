package main

import (
	"fmt"
	"math"
)

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	r float64
}

// methods are special type of function for which we do not need to explicity pass the pointer to
// the struct and a "reciever" is added between the keyword func and the name of the function
// whhich can then be invoked using the . operator for the type in question
// can be invoked with c.area()
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// For Rectangle
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

func main() {
	c := Circle{r: 5}
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	a := Point{x: 0, y: 0}
	b := Point{x: 10, y: 10}
	r := Rectangle{a: a, b: b}
	fmt.Println(r.area())
}
