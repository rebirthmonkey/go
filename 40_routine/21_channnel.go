package main

import "fmt"

func worker(stopCh chan bool) {
	fmt.Println("start goroutine ...")
	stopCh <- true
	fmt.Println("end goroutine...")
}

func main() {
	//stopCh := make(chan bool, 1)
	stopCh := make(chan bool)
	go worker(stopCh)
	fmt.Println("the value in stopCh is: ", <-stopCh)
	fmt.Println("end main...")
}
