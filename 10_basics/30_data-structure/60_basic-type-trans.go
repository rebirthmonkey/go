package main

import (
	"fmt"
)

func main() {
	var a int64 = 3
	var b int32
	b = int32(a)
	fmt.Println("the value of b is: ", b)

	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Println("the mean is: ", mean)
}
