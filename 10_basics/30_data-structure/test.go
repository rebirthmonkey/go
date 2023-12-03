package main

import "fmt"

func changeOrder(data []int) []int {
	var data2 = make([]int, len(data))

	var leftIndex int = 0
	var rightIndex int = len(data) - 1

	for _, value := range data {
		if value%2 == 0 {
			data2[rightIndex] = value
			rightIndex--
		} else {
			data2[leftIndex] = value
			leftIndex++
		}
	}
	return data2
}

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("the new data is: ", changeOrder(data))
}
