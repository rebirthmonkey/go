package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main() {
	fmt.Println("max of 10 and 20 is", max(10, 20))

	i := 10
	fmt.Println("Value of i before increment is : ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is : ", i)
}
