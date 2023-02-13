package main

import (
	"fmt"
)

func main() {
	var a int64 = 3
	var b int32
	b = int32(a)
	fmt.Println("b 的值为：", b)

	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum)/float32(count)
	fmt.Println("mean 的值为: ", mean)
}
