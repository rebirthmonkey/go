package main

import (
	"fmt"
)

//当捕获 panic 时触发此函数
func doPanic() {
	err := recover()
	if err != nil {
		fmt.Println("捕获panic")
	}
}

func main() {
	//注册捕获 panic 函数，必须先注册，若在 panic 之后则无效
	defer doPanic()
	n := 0
	res := 1 / n
	fmt.Println(res) //panic 之后的代码不会执行
}
