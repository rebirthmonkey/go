package main

import (
	"fmt"
)

type student struct {
	rollNo int
	name string
}

func main() {
	stud := student{1, "Johnny"}
	fmt.Println(stud)
	fmt.Println("Student name is:", stud.name)

	fmt.Println(student{rollNo: 2, name: "Ann"})
	fmt.Println(student{name: "Ann", rollNo: 2})
	fmt.Println(student{name: "Alice"})

	pstud := &stud
	fmt.Println("Student name is:", pstud.name)
}
