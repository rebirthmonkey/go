package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("hello")
}

func main() {
	a := Animal{Name: "dog"}
	reflect.ValueOf(&a).MethodByName("Speak").Call([]reflect.Value{})
}
