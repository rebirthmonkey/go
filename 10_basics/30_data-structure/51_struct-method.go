package main

import (
	"fmt"
)

type Rect struct {
	width float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{width: 10, height: 10}
	fmt.Println("Area is: ", r.Area())
	fmt.Println("Perimeter is: ", r.Perimeter())

	ptr := &Rect{width: 10, height: 5}
	fmt.Println("Area is: ", ptr.Area())
	fmt.Println("Perimeter is: ", ptr.Perimeter())
}
