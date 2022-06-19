package main

import (
	"fmt"
)

func main() {
	var v1 int
	v1 = 100 // variable declaration and initialization

	var v2 int = 200 // variable declaration and initialization

	var v3 int // variable declaration without initialization

	v4 := 400 // short variable declaration

	const c1 = 3.14

	fmt.Println("Value stored in variable v1 :: ", v1)
	fmt.Println("Value stored in variable v2 :: ", v2)
	fmt.Println("Value stored in variable v3 :: ", v3)
	fmt.Println("Value stored in variable v4 :: ", v4)
	fmt.Println("Value stored in constant c1 :: ", c1)
}
