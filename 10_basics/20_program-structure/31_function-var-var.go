package main

import "fmt"

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc2(args []int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	var int_slice = []int{1, 2, 3, 4, 5}
	myfunc(int_slice...)
	myfunc2(int_slice)
}
