package main

import (
	"fmt"
)

/*
	遍历切片，对遍历中访问的每个元素使用匿名函数来实现操作。
	用户传入不同的匿名函数体可以实现不同的操作。
*/

// 遍历切片中每个元素，通过给定的函数对每个元素进行操作
func visit(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}

func main() {
	l := []int{1,2,3}
	// 使用匿名函数打印切片的内容
	visit(l, func(value int) {
		fmt.Println(value)
	})
}