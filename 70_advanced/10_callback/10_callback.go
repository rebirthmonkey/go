package main

import (
	"fmt"
)

// 遍历切片中每个元素，通过给定的函数对每个元素进行操作
func visit(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}

func main() {
	l := []int{1, 2, 3}
	// 使用匿名函数打印切片的内容
	visit(l, func(value int) {
		fmt.Println("print the value:", value)
	})
}
