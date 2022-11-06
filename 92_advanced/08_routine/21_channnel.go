package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("start goroutine ...")
	done <- true
	fmt.Println("end goroutine...")
}

func main() {
	//done := make(chan bool, 1)
	done := make(chan bool)
	go worker(done)
	<-done
}
