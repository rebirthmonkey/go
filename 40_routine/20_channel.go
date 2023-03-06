package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second) // 延时1秒
		}
	}()

	time.Sleep(time.Second * 2)

	msg := <-messages
	fmt.Println("the current value in the messages is:", msg)

	time.Sleep(time.Second * 3)
}
