package main

import (
	"fmt"
	"time"
)

func running() {
	var times int

	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second) // 延时1秒
	}
}

func main() {
	go running()  // 并发执行程序

	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
	println("the input is: ", input)
}

