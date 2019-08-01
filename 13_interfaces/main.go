package main

import (
	"fmt"
	"math"
)

// The Shape Interface
type Shape interface {
	area() float64
}

// The Circle Struct
type Circle struct {
	x,y, radius float64
}

// The Rectangle Struct
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 1, y: 1, radius:5}
	rectangle := Rectangle{height: 10,width: 15}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
