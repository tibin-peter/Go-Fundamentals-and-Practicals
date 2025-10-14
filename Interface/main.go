package main

import (
	"fmt"
	"math"
)

// define interface methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// for rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// for circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f,Perimeter:%2f\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 5}

	PrintShapeInfo(c)
	PrintShapeInfo(r)
}
