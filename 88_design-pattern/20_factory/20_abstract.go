package main

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// 此处返回Person interface
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age: age,
	}
}

func main(){
	p := NewPerson("Tom", 30)
	p.Greet()
}
