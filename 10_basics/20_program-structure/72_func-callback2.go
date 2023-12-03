package main

type callback2 func([]int)

func treatIntSlice(in []int, cb callback2) {
	cb(in)
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}

	treatIntSlice(intSlice, func(intSlice []int) {
		for _, val := range intSlice {
			println("the 1st callback func: ", val)
		}
	})

	treatIntSlice(intSlice, func(intSlice []int) {
		for _, val := range intSlice {
			println("the 2nd callback func: ", val)
		}
	})

	return
}
