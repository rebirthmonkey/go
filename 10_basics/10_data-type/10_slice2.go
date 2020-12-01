package main

import "fmt"

func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

func main() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		PrintSlice (s)
	}
}

