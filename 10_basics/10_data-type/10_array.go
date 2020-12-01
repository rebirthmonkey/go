package main

import "fmt"

func main() {
	var arr1 [10]int
	fmt.Println("arr1", arr1)
	for i := 0; i < 10; i++ {
		arr1[i] = i
	}
	fmt.Println("arr1", arr1)
	fmt.Println("Length of arr1", len(arr1))

	arr2 := [3]int{9,8,7}
	fmt.Println("arr2", arr2)
}
