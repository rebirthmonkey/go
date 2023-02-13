package main

import "fmt"

func main() {

	c := make(chan string)

	go func() {
		fmt.Println("Go ...!")
		c <- "aaa"
		c <- "bbb"
		c <- "ccc"
		c <- "ddd"
		close(c)
	}()

	msg := <-c
	fmt.Println(msg)

	for v := range c {
		fmt.Println("rang:", v)
	}

	fmt.Println("Finish ...!")
}
