package main

import (
	"fmt"
)

// 函数 return 前执行已注册 defer
func f1() {
	defer fmt.Println("f1()：return 前执行 defer")
	return
}

// 函数执行到最后执行已注册 defer
func f2() {
	defer fmt.Println("f2()：函数执行到最后执行已注册 defer")
	fmt.Println("f2()：函数执行")
}

// panic 前执行已注册的 defer
func f3() {
	defer fmt.Println("f3()：panic 前注册 defer") // 已注册 defer
	panic("f3()：panic 中")
	defer fmt.Println("f3()：panic 后未注册 defer") // 未注册 defer
}

func main(){
	f1()
	f2()
	f3()
}
