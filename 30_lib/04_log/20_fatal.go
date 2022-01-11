package main

import (
	"fmt"
	"log"
)

func deferFatal(){
	defer func() {
		fmt.Println("defer.....")
	}()
	log.Fatalln("test for defer Fatal")
}

func main() {
	deferFatal()
}
