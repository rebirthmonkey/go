package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

// 定义命令行参数对应的变量
var cliName = pflag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = pflag.IntP("age", "a", 22, "Input Your Age")
var cliGender = pflag.StringP("gender", "g", "male", "Input Your Gender")
var cliOK = pflag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = pflag.StringP("des-detail", "d", "", "Input Description")

func main() {

	pflag.Parse() // 把用户传递的命令行参数解析为对应变量的值

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}
