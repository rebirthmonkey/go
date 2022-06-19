package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"strings"
)

// 定义命令行参数对应的变量
var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = flag.IntP("age", "a",22, "Input Your Age")
var cliGender = flag.StringP("gender", "g","male", "Input Your Gender")
var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")


func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}


func main() {
	// 设置"名字标准化函数"
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// 为age参数设置NoOptDefVal
	flag.Lookup("age").NoOptDefVal = "25"

	// 把badflag参数标记为即将废弃的，请用户使用des-detail参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把badflag参数的shorthand标记为即将废弃的
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数badflag
	flag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}
