package main

import "fmt"

func main() {
	var a interface{} = 100
	if b, ok := a.(int); ok {
		fmt.Println("the assert translated value is:", b)
	}
}
