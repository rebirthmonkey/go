package main

import (
	"fmt"
)

var block string = "package"

func main() {
	block := "function"
	{ // 这边额外添加了一个 block
		block := "inner"
		fmt.Println("The block is", block)
	}
	fmt.Println("The block is", block)
}
