package main

import (
	"os"
	"text/template"
)

func main() {
	// 模板定义
	text := "My name is {{ . }}\n"

	// 解析模板
	tmpl1, _ := template.New("test").Parse(text)

	// 数据驱动模板
	data := "jack"
	_ = tmpl1.Execute(os.Stdout, data)


	// 数据结构驱动模板
	stu := struct{Name string; ID int}{Name: "hello", ID: 11}

	// 创建模板对象, parse关联模板
	tmpl2, _ := template.New("test").Parse("{{.Name}} ID is {{ .ID }}\n")

	// 渲染stu为动态数据, 标准输出到终端
	_ = tmpl2.Execute(os.Stdout, stu)
}


