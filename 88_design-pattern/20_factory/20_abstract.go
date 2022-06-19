package main

import "fmt"

type Person2 interface {
	Greet()
}

type person2 struct {
	name string
	age int
}

func (p person2) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// 此处返回Person2 interface
func NewPerson2(name string, age int) Person2 {
	return person2{
		name: name,
		age: age,
	}
}

func main(){
	p := NewPerson2("Tom", 30)
	p.Greet()
}
