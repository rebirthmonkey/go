package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
	println("the input is: ", input)
}


