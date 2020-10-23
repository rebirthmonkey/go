package main

import "fmt"

func main() {
	s0 := []int{0,1,2,3,4}
	fmt.Printf("The value of s1: %d\n", s0)
	fmt.Printf("The length of s1: %d\n", len(s0))
	fmt.Printf("The capacity of s1: %d\n", cap(s0))
	s1 := make([]int, 5)
	fmt.Printf("The value of s1: %d\n", s1)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	s2 := make([]int, 5, 8)
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
}
