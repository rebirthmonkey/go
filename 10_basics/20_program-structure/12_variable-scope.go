package main

import (
	"fmt"
)

var block string = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Println("The block is", block)
	}
	fmt.Println("The block is", block)
}
