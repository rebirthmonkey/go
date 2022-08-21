package main

import (
	"fmt"
	"hungry"
	"lazy"
)

func main() {
	fmt.Println("the mode of hungry is", hungry.GetInsOr())
	fmt.Println("the mode of layz is", lazy.GetInsOr())
}
