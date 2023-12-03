package main

import (
	"fmt"
	"time"
)

var c = make(chan struct{})
var f = make(chan struct{})
var d = make(chan struct{})

func cat() {
	<-d
	fmt.Println("cat")
	c <- struct{}{}
}

func fish() {
	<-c
	fmt.Println("fish")
	f <- struct{}{}
}

func dog() {
	<-f
	fmt.Println("dog")
	d <- struct{}{}
}

func main() {
	for i := 0; i < 10; i++ {
		go cat()
		go fish()
		go dog()
	}
	d <- struct{}{}

	time.Sleep(10 * time.Second)

	fmt.Println("finish all")
}
