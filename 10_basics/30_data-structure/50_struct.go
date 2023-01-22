package main

import (
	"fmt"
)

type student struct {
	rollNo int
	name   string
}

type graduate struct {
	level int
	*student
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

	grad := graduate{
		level:   8,
		student: &stud,
	}
	fmt.Println("Graduate student name is:", grad.name)
}
