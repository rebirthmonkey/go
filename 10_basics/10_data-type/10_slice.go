package main

import "fmt"


func main() {
	var sli1 []int
	fmt.Println("sli1", sli1)
	for i := 1; i <= 17; i++ {
		sli1 = append(sli1, i)
	}
	fmt.Println("sli1", sli1)

	sli2 := []int{11,12,13}
	sli2 = append(sli2, 14)
	fmt.Println("sli2", sli2)

	sli3 := make([]int, 10)
	sli3 = append(sli3, 22)
	fmt.Println("sli3", sli3)
}

