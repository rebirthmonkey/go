package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is:", sum)

	sum2 := 0
	for _, val := range numbers {
		sum2 += val
	}
	fmt.Println("Sum2 is:", sum2)
}
