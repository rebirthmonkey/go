package main

import (
	"fmt"
)

type fileError struct {
	s string
}

// 实现了Error()函数就等于实现了error interface
func (fe *fileError) Error() string {
	return fe.s
}

//模拟一个错误
func openFile() ([]byte, error) {
	return nil, &fileError{"自定义文件错误"}
}

func main() {
	content, err := openFile()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}
}

