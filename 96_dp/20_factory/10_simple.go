package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

func main() {
	p := NewPerson("Tom", 30)
	p.Greet()
}
