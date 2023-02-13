package main

import (
	"fmt"
	"time"
)

func running() {
	var times int

	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	var input string
	fmt.Scanln(&input)
	println("the input is: ", input)
}
