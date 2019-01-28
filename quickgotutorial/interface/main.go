package main

import (
	"fmt"
	"math"
)

//Define Interface

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
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
	circle := Circle{3, 4, 5}
	rectangle := Rectangle{3, 4}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}
