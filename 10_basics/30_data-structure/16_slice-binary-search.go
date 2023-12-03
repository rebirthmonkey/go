package main

func BinarySearch(data []int, value int) bool {
	var low int = 0
	var high int = len(data) - 1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		println("mid: ", mid)
		if data[mid] == value {
			return true
		} else {
			if data[mid] > value {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

func main() {
	var d = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	println("the binary search result is: ", BinarySearch(d, 20))
	println("the binary search result is: ", BinarySearch(d, 13))
}
