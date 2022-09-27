package main

import (
	"fmt"
	"math"
)

type Shape2 interface {
	Area() float64
	Perimeter() float64
}

type Rec2 struct {
	width  float64
	height float64
}

func (r Rec2) Area() float64 {
	return r.width * r.height
}

func (r Rec2) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle2 struct {
	radius float64
}

func (c Circle2) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle2) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea2(shapes ...Shape2) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter2(shapes ...Shape2) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}

func main() {
	r := Rec2{width: 10, height: 10}
	c := Circle2{radius: 10}
	// 任何实现了interface的struct都可以作为接口"基类"的子类的实例参数传给函数
	fmt.Println("Total Area is: ", TotalArea2(r, c))
	fmt.Println("Total Perimeter is: ", TotalPerimeter2(r, c))
}
