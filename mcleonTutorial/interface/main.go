package main

import (
	"fmt"
	"math"
)

type Square struct {
	size float64
}
type Circle struct {
	size float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.size * s.size
}
func (c Circle) area() float64 {
	return math.Pi * c.size * c.size
}

func info(s Shape) {
	fmt.Println(s.area())
	fmt.Println(s)
}
func main() {
	s := Square{size: 10}
	c := Circle{size: 10}
	info(s)
	info(c)
	fmt.Println("Ya")
}
