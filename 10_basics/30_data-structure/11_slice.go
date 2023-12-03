package main

import (
	"fmt"
)

func main() {
	var sli1 []int
	fmt.Println("sli1 is:", sli1)
	for i := 0; i <= 7; i++ {
		sli1 = append(sli1, i)
	}
	fmt.Println("sli1 is:", sli1)
	fmt.Println("Length of sli1:", len(sli1))
	fmt.Println("Capacity of sli1", cap(sli1))

	sli2 := []int{11, 12, 13}
	sli2 = append(sli2, 14)
	fmt.Println("sli2 is:", sli2)

	sli3 := make([]int, 3)
	sli3 = append(sli3, 22)
	fmt.Println("sli3 is:", sli3)
}
