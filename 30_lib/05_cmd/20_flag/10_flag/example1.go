package main

import (
	"flag"
	"fmt"
)

func main(){
	// 把flat参数注册到变量指针port上
	port := flag.Int("p", 8080, "server Port")
	// 解析参数
	flag.Parse()
	fmt.Println("server port:", *port)
}
