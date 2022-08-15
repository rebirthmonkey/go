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

	var s2 string
	for _, p := range os.Args[1:]{ // 否则Args[0]也会展示
		s2 += sep + p
	}
	fmt.Println(s2)
}
