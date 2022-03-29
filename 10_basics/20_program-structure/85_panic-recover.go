package main

import (
	"fmt"
)

//当捕获 panic 时触发此函数
func doPanic() {
	err := recover()
	if err != nil {
		fmt.Println("defer函数捕获panic")
	}
}

func main() {
	defer doPanic() // 注册捕获panic的defer函数，必须先注册，若在panic之后则无效
	n := 0
	res := 1 / n
	fmt.Println(res) //panic之后的代码不会执行
}
