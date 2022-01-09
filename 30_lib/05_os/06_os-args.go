package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = " "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}

// go run 06_os-args.go a b c d