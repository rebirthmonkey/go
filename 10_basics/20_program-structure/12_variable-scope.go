package main

import (
	"fmt"
)

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Println("The block is", block)
	}
	fmt.Println("The block is", block)
}
