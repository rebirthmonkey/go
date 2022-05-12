package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	
	for { // 构建一个无限循环
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second) // 延时1秒
	}
}

func main() {
	go running()

	var input string
	fmt.Scanln(&input)
	println("the input is: ", input)
}
