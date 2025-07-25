package main

import (
	"fmt"
	"math"
)
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func interfaces() { 
	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is ", getArea(rect))
	fmt.Println("Area of circle is ", getArea(circ))
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}