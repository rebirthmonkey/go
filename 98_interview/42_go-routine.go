package main

import (
	"fmt"
	"time"
)

var a = make(chan struct{})
var b = make(chan struct{})

func printA() {

	for i := 0; i < 10; i++ {
		<-a
		println("A")
		b <- struct{}{}
	}
}

func printB() {

	for i := 0; i < 10; i++ {
		<-b
		println("B")
		a <- struct{}{}
	}
}

func main() {

	go printA()
	go printB()
	a <- struct{}{}

	time.Sleep(5 * time.Second)
	fmt.Println("finish all")
}
