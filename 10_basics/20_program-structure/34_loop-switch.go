package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something else")
	}

	switch i {
	case 1, 2, 3:
		fmt.Println("one, two or three")
	default:
		fmt.Println("something else")
	}
}
