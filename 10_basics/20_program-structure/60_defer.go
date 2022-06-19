package main

import (
	"fmt"
)

// 函数return前执行已注册defer
func f1() {
	defer fmt.Println("f1()：return前执行defer")
	return
}

// 函数执行到最后执行已注册defer
func f2() {
	defer fmt.Println("f2()：函数执行到最后执行已注册defer")
	fmt.Println("f2()：函数执行")
}

// panic前执行已注册的defer
func f3() {
	defer fmt.Println("f3()：panic前注册defer") // 已注册defer
	panic("f3()：panic中")
	defer fmt.Println("f3()：panic后未注册defer") // 未注册defer
}

func main(){
	f1()
	f2()
	f3()
}
