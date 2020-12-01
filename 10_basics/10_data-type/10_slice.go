package main

import "fmt"

<<<<<<< HEAD

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

=======
func main() {
	s0 := []int{0,1,2,3,4}
	fmt.Printf("The value of s1: %d\n", s0)
	fmt.Printf("The length of s1: %d\n", len(s0))
	fmt.Printf("The capacity of s1: %d\n", cap(s0))
	s1 := make([]int, 5)
	fmt.Printf("The value of s1: %d\n", s1)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	s2 := make([]int, 5, 8)
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
}
>>>>>>> de31d1d040ee13ebba7944684f3bebb3f7196274
