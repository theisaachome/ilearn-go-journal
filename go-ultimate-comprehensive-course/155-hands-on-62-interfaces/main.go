package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type Circle struct {
	Radius float64
}

func (s square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type shape interface {
	Area() float64
}

func info(s shape) {
	fmt.Println("Area:", s.Area())
}
func main() {
	square := square{length: 5, width: 6}
	info(square)

	circle := Circle{Radius: 5}
	info(circle)
}
