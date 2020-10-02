package main

func SequentialSearch(data []int, value int) bool{
	for _, _value := range data{
		if _value == value {
			return true
		}
	}
	return false
}

func main(){
	var d = []int{1,2,3,4,5}
	println("The sequential search value is: ", SequentialSearch(d, 3))
}

