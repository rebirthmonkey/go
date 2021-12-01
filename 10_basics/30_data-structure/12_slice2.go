package main

import (
	"fmt"
)

func printSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

func main() {
	var sli []int
	for i := 0; i <= 5; i++ {
		sli = append(sli, i)
		printSlice (sli)
	}
}

