package main

import (
	"fmt"
	"log"
)

func deferPanic(){
	defer func() {
		fmt.Println("defer1.....")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("logPanic: test")
	defer func() {
		fmt.Println("defer2.....")
	}()
}

func main() {
	deferPanic()
}
