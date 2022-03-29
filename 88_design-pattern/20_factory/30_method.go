package main

import "fmt"

type Person3 struct {
	name string
	age int
}

func NewPersonFactory(age int) func(name string) Person3 {
	return func(name string) Person3 {
		return Person3{
			name: name,
			age: age,
		}
	}
}

func main() {
	babyFactory := NewPersonFactory(1)
	baby := babyFactory("john")
	fmt.Println("the new baby created by the BabyFactory is", baby)

	teenagerFactory := NewPersonFactory(16)
	teen := teenagerFactory("jill")
	fmt.Println("the new teenager create by the TeenagerFactory is:", teen)
}
