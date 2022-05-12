package main

import (
	"fmt"
	"time"
)

func strWorker(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with strWorker...")
	ch <- "str"
}

func intWorker(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with intWorker...")
	ch <- 1
}

func main() {
	chStr := make(chan string)
	chInt := make(chan int)

	go strWorker(chStr)
	go intWorker(chInt)

	for i := 0; i < 2; i++ {
		select {
		case <-chStr:
			fmt.Println("get value from strWorker")

		case <-chInt:
			fmt.Println("get value from intWorker")
		}
	}
}
