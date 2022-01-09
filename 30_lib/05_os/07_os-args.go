package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}

// go run 07_os-args.go a b c d