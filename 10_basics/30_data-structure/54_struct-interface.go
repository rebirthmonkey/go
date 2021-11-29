package main

import (
	"fmt"
	"math"
)

type Rec struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rec) Area() float64 {
	return r.width * r.height
}

func (r Rec) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}

func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println("Total Area is: ", TotalArea(r, c))
	fmt.Println("Total Perimeter is: ", TotalPerimeter(r, c))
}
