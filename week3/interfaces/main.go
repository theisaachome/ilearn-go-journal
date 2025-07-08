package main

import (
	"fmt"
	"math"
)


type Shape interface{
	Area() float64
	Perimeter() float64
}
// interface composition
type Measurable interface{

	Perimeter() float64
}

type Geometry interface{
	Shape
	Measurable
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func describeShape(g Geometry){
	fmt.Printf("Shape Area: %f, Perimeter: %f\n", g.Area(), g.Perimeter())
}
func calculateArea(s Shape) float64 {
	return s.Area()
}

func main(){
	 c :=Circle{Radius: 5}
	 println("Circle Area:", calculateArea(c))

	 // rectangle example

	  r := Rectangle{Width: 4, Height: 6}
	  println("Rectangle Area:", calculateArea(r))

	  describeShape(r)

	  myVariable :=interface{}(42) // using empty interface to hold any type
	  describeVariable(myVariable)


}

func describeVariable(v interface{}){
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}
