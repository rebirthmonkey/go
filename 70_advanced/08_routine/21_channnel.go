package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
