package main

func SumSlice(data []int) int {
	total := 0
	for index := 0; index < len(data); index++ {
		total += data[index]
	}
	return total
}

func main() {
	var l = []int{1, 2, 3, 4, 5}
	println("the Sum Array is: ", SumSlice(l))
}
