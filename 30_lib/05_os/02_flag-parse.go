package main

import (
	"flag"
	"fmt"
)

func main(){
	// 注册参数
	port := flag.Int("p", 8080, "server Port")
	// 解析参数
	flag.Parse()
	fmt.Println("server port:", *port)
}
