package main

import (
	"fmt"
)

type Animal2 interface {
	Speak() string
}

type Dog2 struct {
}

func (d Dog2) Speak() string {
	return "Woof!"
}

type Cat2 struct {
}

func (c *Cat2) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal2{Dog2{}, &Cat2{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
