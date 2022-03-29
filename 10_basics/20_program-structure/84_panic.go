package main

import (
	"fmt"
)

func main() {
	n := 0
	res := 1 / n
	fmt.Println(res) // panic 之后的代码不会执行
}
