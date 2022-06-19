package main

func SumSlice(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func main(){
	var l = []int{1,2,3,4,5}
	println("the Sum Array is: ", SumSlice(l))
}
